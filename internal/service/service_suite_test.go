package service_test

import (
	mock_config "github.com/barisaskaleli/lightweight-bank/config/mocks"
	mock_repository "github.com/barisaskaleli/lightweight-bank/internal/repository/mocks"
	"github.com/barisaskaleli/lightweight-bank/internal/service"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Service Suite")
}

var (
	ctrl           *gomock.Controller
	mockConfig     *mock_config.MockIConfig
	mockLogger     *zap.SugaredLogger
	mockRepository *mock_repository.MockIRepository
	svc            service.IService
)

var _ = BeforeSuite(func() {
	ctrl = gomock.NewController(GinkgoT())
	mockConfig = mock_config.NewMockIConfig(ctrl)
	mockLogger = zap.NewNop().Sugar()
	mockRepository = mock_repository.NewMockIRepository(ctrl)
	svc = service.BuildService(mockConfig, mockLogger, mockRepository)
})

var _ = AfterSuite(func() {
	ctrl.Finish()
})
