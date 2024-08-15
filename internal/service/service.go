package service

import (
	"errors"
	"github.com/barisaskaleli/lightweight-bank/config"
	"github.com/barisaskaleli/lightweight-bank/internal/repository"
	"github.com/barisaskaleli/lightweight-bank/internal/resource/model"
	"github.com/barisaskaleli/lightweight-bank/internal/resource/request"
	"github.com/barisaskaleli/lightweight-bank/internal/resource/response"
	"github.com/barisaskaleli/lightweight-bank/internal/util"
	"github.com/golang-jwt/jwt/v5"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sync"
	"time"
)

type IService interface {
	Register(req request.RegisterRequest) (response.RegisterResponse, error)
	Login(req request.LoginRequest) (response.LoginResponse, error)
	Transfer(req request.TransferRequest) (response.TransferResponse, error)
}

type service struct {
	config     config.IConfig
	logger     *zap.SugaredLogger
	repository repository.IRepository
	mutex      sync.Mutex
}

func BuildService(config config.IConfig, logger *zap.SugaredLogger, repository repository.IRepository) IService {
	return &service{
		config:     config,
		logger:     logger,
		repository: repository,
	}
}

func (s *service) Register(req request.RegisterRequest) (response.RegisterResponse, error) {
	var res response.RegisterResponse

	accNumber := util.GenerateAccountNumber()

	user, err := s.repository.Register(req, accNumber)

	if err != nil {
		s.logger.Errorf("Error while registering user: %v", err)
		return res, err
	}

	res.ID = user.ID
	res.AccountNumber = user.AccountNumber
	res.Name = user.Name
	res.Surname = user.Surname
	res.Email = user.Email
	res.Balance = user.Balance

	return res, nil
}

func (s *service) Login(req request.LoginRequest) (response.LoginResponse, error) {
	var res response.LoginResponse

	user, err := s.repository.Login(req)

	if err != nil {
		s.logger.Errorf("Error while logging in: %v", err)
		return res, err
	}

	claims := jwt.MapClaims{
		"email":          user.Email,
		"account_number": user.AccountNumber,
		"exp":            time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(s.config.Server().JWTSecret))

	if err != nil {
		s.logger.Errorf("Error while signing token: %v", err)
		return res, err
	}

	res.Token = signedToken

	return res, nil
}

func (s *service) Transfer(req request.TransferRequest) (response.TransferResponse, error) {
	sender, receiver, err := s.getBothUsers(req.Sender, req.Receiver)
	if err != nil {
		return s.createTransferErrorResponse(err.Error(), err)
	}

	senderBalance, receiverBalance, err := s.validateTransfer(sender.Balance, receiver.Balance, req.Amount)
	if err != nil {
		return s.createTransferErrorResponse(err.Error(), err)
	}

	tx := s.repository.GetDatabase().Database().Begin()

	err = s.updateBalances(tx, sender.AccountNumber, receiver.AccountNumber, senderBalance, receiverBalance)
	if err != nil {
		errMsg := "Error while updating balances"
		s.logger.Errorf("%s: %v", errMsg, err)
		tx.Rollback()
		return s.createTransferErrorResponse(errMsg, err)
	}

	err = s.repository.AddTransaction(tx, sender, receiver, req.Amount, s.config.Service().TransferFee)
	if err != nil {
		errMsg := "Error while adding transaction"
		s.logger.Errorf("%s: %v", errMsg, err)
		tx.Rollback()
		return s.createTransferErrorResponse(errMsg, err)
	}

	err = tx.Commit().Error
	if err != nil {
		errMsg := "Error while committing transaction"
		s.logger.Errorf("%s: %v", errMsg, err)
		return s.createTransferErrorResponse(errMsg, err)
	}

	if s.config.Service().SMSInfoEnabled == true {
		go s.sendSMS()
	}

	return s.createTransferSuccessResponse(senderBalance, receiverBalance), nil
}

func (s *service) getBothUsers(senderAcc, receiverAcc string) (sender, receiver model.User, err error) {
	sender, err = s.repository.GetByAccountNumber(senderAcc)

	if err != nil {
		s.logger.Errorf("Error while getting sender: %v", err)
		return sender, receiver, err
	}

	receiver, err = s.repository.GetByAccountNumber(receiverAcc)

	if err != nil {
		s.logger.Errorf("Error while getting receiver: %v", err)
		return sender, receiver, err
	}

	return sender, receiver, nil
}

func (s *service) validateTransfer(sBalance, rBalance, amount float64) (float64, float64, error) {
	senderBalance := decimal.NewFromFloat(sBalance)
	receiverBalance := decimal.NewFromFloat(rBalance)
	senderAmount := decimal.NewFromFloat(amount)

	if s.config.Service().TransferFeeEnabled == true {
		transferFee := decimal.NewFromFloat(s.config.Service().TransferFee)
		senderAmount = senderAmount.Add(transferFee)
	}

	if senderBalance.LessThan(senderAmount) {
		return senderBalance.InexactFloat64(), receiverBalance.InexactFloat64(), errors.New("insufficient balance")
	}

	exactAmount := decimal.NewFromFloat(amount)

	senderBalance = senderBalance.Sub(senderAmount)
	receiverBalance = receiverBalance.Add(exactAmount)

	return senderBalance.InexactFloat64(), receiverBalance.InexactFloat64(), nil
}

func (s *service) updateBalances(tx *gorm.DB, senderAcc, receiverAcc string, senderBalance, receiverBalance float64) error {
	err := s.repository.UpdateBalance(tx, senderAcc, senderBalance)
	if err != nil {
		return err
	}

	err = s.repository.UpdateBalance(tx, receiverAcc, receiverBalance)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) sendSMS() {
	s.logger.Info("Sending SMS")
}

func (s *service) createTransferSuccessResponse(sender, receiver float64) response.TransferResponse {
	return response.TransferResponse{
		Status:          true,
		Message:         "Transfer completed successfully",
		SenderBalance:   &sender,
		ReceiverBalance: &receiver,
		Fee:             s.config.Service().TransferFee,
	}
}

func (s *service) createTransferErrorResponse(message string, err error) (response.TransferResponse, error) {
	return response.TransferResponse{
		Status:  false,
		Message: message,
	}, err
}
