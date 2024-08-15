package repository

import (
	"github.com/barisaskaleli/lightweight-bank/config"
	"github.com/barisaskaleli/lightweight-bank/internal/resource/model"
	"github.com/barisaskaleli/lightweight-bank/internal/resource/request"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IRepository interface {
	GetDatabase() config.IMysqlInstance
	Register(request request.RegisterRequest, accountNumber string) (model.User, error)
	Login(request request.LoginRequest) (model.User, error)
	GetByAccountNumber(accountNumber string) (model.User, error)
	UpdateBalance(tx *gorm.DB, accountNumber string, amount float64) error
	AddTransaction(tx *gorm.DB, sender, receiver model.User, amount, fee float64) error
}

type repository struct {
	db     config.IMysqlInstance
	cfg    config.IConfig
	logger *zap.SugaredLogger
}

func BuildRepository(db config.IMysqlInstance, cfg config.IConfig, logger *zap.SugaredLogger) IRepository {
	return &repository{
		db:     db,
		cfg:    cfg,
		logger: logger,
	}
}

func (r *repository) GetDatabase() config.IMysqlInstance {
	return r.db
}

func (r *repository) Register(request request.RegisterRequest, accountNumber string) (model.User, error) {
	user := model.User{
		AccountNumber: accountNumber,
		Name:          request.Name,
		Surname:       request.Surname,
		Email:         request.Email,
		Password:      request.Password,
		Balance:       request.Balance,
	}

	err := r.db.Database().Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Login(request request.LoginRequest) (model.User, error) {
	var user model.User

	err := r.db.Database().Where("email = ? AND password = ?", request.Email, request.Password).First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) GetByAccountNumber(accountNumber string) (model.User, error) {
	var user model.User

	err := r.db.Database().Where("account_number = ?", accountNumber).First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) UpdateBalance(tx *gorm.DB, accountNumber string, amount float64) error {
	return tx.Model(&model.User{}).Where("account_number = ?", accountNumber).Update("balance", amount).Error
}

func (r *repository) AddTransaction(tx *gorm.DB, sender, receiver model.User, amount, fee float64) error {
	transaction := model.Transaction{
		Sender:   sender,
		Receiver: receiver,
		Amount:   amount,
		Fee:      fee,
		Type:     "transfer",
	}

	return tx.Create(&transaction).Error
}
