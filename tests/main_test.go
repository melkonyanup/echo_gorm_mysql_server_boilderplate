package tests

import (
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"myapp/internal/app"
	"myapp/internal/delivery/http/v1"
	"myapp/internal/repository"
	"myapp/internal/service"
	"myapp/pkg/auth"
	"myapp/pkg/common_services"
	"myapp/pkg/database"
	"os"
	"testing"
)

type APITestSuite struct {
	suite.Suite

	db           *gorm.DB
	handler      *v1.Handler
	tokenManager *auth.TokenManager
	hasher       common_services.Hasher
}

func TestAPISuite(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	suite.Run(t, new(APITestSuite))
}

func (s *APITestSuite) SetupSuite() {
	err := app.InitApp()
	if err != nil {
		s.FailNow("Failed to initialize application", err)
	}

	s.db, err = database.InitDatabase()
	if err != nil {
		s.FailNow("Failed to connect to db", err)
	}

	repos := repository.NewRepository(s.db)
	services := service.NewService(repos)
	s.handler = v1.NewHandler(services)
	s.hasher = &common_services.AppHasher{}
	s.tokenManager = &auth.TokenManager{
		SecretKey:       os.Getenv("ACCESS_TOKEN_SECRET_KEY"),
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}
}

func (s *APITestSuite) TearDownSuite() {
	sqlDB, _ := s.db.DB()
	sqlDB.Close()
}

func TestMain(m *testing.M) {
	rc := m.Run()
	os.Exit(rc)
}
