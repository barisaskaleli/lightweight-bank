package service_test

import (
	"errors"
	"github.com/barisaskaleli/lightweight-bank/config"
	"github.com/barisaskaleli/lightweight-bank/internal/resource/model"
	"github.com/barisaskaleli/lightweight-bank/internal/resource/request"
	"github.com/barisaskaleli/lightweight-bank/internal/resource/response"
	"github.com/golang-jwt/jwt/v5"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

var _ = Describe("Service", func() {
	Context("Login", func() {
		It("should return a valid token for a successful login", func() {
			req := request.LoginRequest{Email: "test@example.com", Password: "password"}
			user := model.User{Email: "test@example.com", AccountNumber: "123456789"}

			mockRepository.EXPECT().Login(req).Return(user, nil)
			mockConfig.EXPECT().Server().Return(config.ServerConfig{JWTSecret: "secret"})

			res, err := svc.Login(req)

			Expect(err).To(BeNil())
			Expect(res.Token).ToNot(BeEmpty())

			token, err := jwt.Parse(res.Token, func(token *jwt.Token) (interface{}, error) {
				return []byte("secret"), nil
			})
			Expect(err).To(BeNil())
			Expect(token.Valid).To(BeTrue())
		})

		It("should return an error if login fails", func() {
			req := request.LoginRequest{Email: "test@example.com", Password: "password"}

			mockRepository.EXPECT().Login(req).Return(model.User{}, errors.New("login failed"))

			res, err := svc.Login(req)

			Expect(err).ToNot(BeNil())
			Expect(res.Token).To(BeEmpty())
		})
	})

	Context("Register", func() {
		It("should return an error if registration fails", func() {
			var res response.RegisterResponse
			req := request.RegisterRequest{
				Name:     "John",
				Surname:  "Doe",
				Email:    "",
				Password: "1234",
				Balance:  100,
			}

			mockRepository.EXPECT().Register(req, gomock.Any()).Return(model.User{}, errors.New("register failed")).AnyTimes()

			res, err := svc.Register(req)

			Expect(err).ToNot(BeNil())
			Expect(res).To(Equal(res))
		})

		It("should return success", func() {
			req := request.RegisterRequest{
				Name:     "John",
				Surname:  "Doe",
				Email:    "johndoe@gmail.com",
				Password: "1234",
				Balance:  100,
			}

			accNumber := "123456"

			mockRepository.EXPECT().Register(req, gomock.Any()).Return(model.User{
				ID:            1,
				AccountNumber: accNumber,
				Name:          req.Name,
				Surname:       req.Surname,
				Email:         req.Email,
				Password:      req.Password,
				Balance:       req.Balance,
			}, nil).AnyTimes()

			res, err := svc.Register(req)

			Expect(err).To(BeNil())
			Expect(res.Name).To(Equal(req.Name))
			Expect(res.Surname).To(Equal(req.Surname))
			Expect(res.Email).To(Equal(req.Email))
			Expect(res.AccountNumber).To(Equal(accNumber))
			Expect(res.Balance).To(Equal(req.Balance))
		})
	})

	Context("Transfer", func() {
		It("should return error if sender has insufficient balance", func() {
			req := request.TransferRequest{Sender: "123456", Receiver: "654321", Amount: 150.0}
			sender := model.User{AccountNumber: "123456", Balance: 100.0}
			receiver := model.User{AccountNumber: "654321", Balance: 50.0}

			mockRepository.EXPECT().GetByAccountNumber("123456").Return(sender, nil)
			mockRepository.EXPECT().GetByAccountNumber("654321").Return(receiver, nil)

			mockConfig.EXPECT().Service().Return(config.ServiceConfig{TransferFeeEnabled: true, TransferFee: 5}).AnyTimes()

			res, err := svc.Transfer(req)

			Expect(err).ToNot(BeNil())
			Expect(res.Status).To(BeFalse())
			Expect(res.Message).To(Equal("insufficient balance"))
		})
	})
})
