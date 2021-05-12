package service

import (
	"myapp/internal/models"
	"myapp/internal/repository"
	"myapp/internal/shared/payloads"
	"myapp/pkg/auth"
	"myapp/pkg/common_services"
	"os"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) SignUp(payload *payloads.SignUpPayload) (error) {
	hasher := common_services.AppHasher{}
	hashedPass, err := hasher.HashPassword(payload.Password)
	if err != nil {
		return err
	}

	user := &models.User{
		FirstName: payload.FirstName, LastName: payload.LastName, Email: payload.Email, Password: hashedPass,
	}
	err = s.repo.CreateUserRecord(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) SignIn(payload *payloads.SignInPayload) (string, error) {
	user, err := s.repo.FindUserByEmail(&payload.Email)
	if err != nil {
		return "", err
	}

	hasher := common_services.AppHasher{}
	err = hasher.CheckPassword(user.Password, payload.Password)
	if err != nil {
		return "", err
	}

	tokenManager := auth.TokenManager{
		SecretKey:       os.Getenv("ACCESS_TOKEN_SECRET_KEY"),
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}
	signedToken, err := tokenManager.GenerateToken(user.Email, user.ID)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *UserService) GetUserProfile(email string) (*models.User, error) {
	return s.repo.FindUserByEmail(&email)
}
