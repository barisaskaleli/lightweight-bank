package service

import (
	"github.com/barisaskaleli/lightweight-bank/config"
	"github.com/barisaskaleli/lightweight-bank/internal/repository"
	"github.com/barisaskaleli/lightweight-bank/internal/resource/request"
	"github.com/barisaskaleli/lightweight-bank/internal/resource/response"
	"github.com/barisaskaleli/lightweight-bank/internal/util"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"time"
)

type IService interface {
	Register(req request.RegisterRequest) (response.RegisterResponse, error)
	Login(req request.LoginRequest) (response.LoginResponse, error)
}

type service struct {
	config     config.IConfig
	logger     *zap.SugaredLogger
	repository repository.IRepository
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
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
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
