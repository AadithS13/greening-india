package service

import (
	"errors"

	"greening-india/internal/model"
	"greening-india/internal/repository"
	"greening-india/internal/utils"

	"gorm.io/gorm"
)

type AuthService struct {
	repo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(name, email, password string) error {
	_, err := s.repo.GetByEmail(email)

	if err == nil {
		return errors.New("email already exists")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	hash, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := model.User{
		Name:     name,
		Email:    email,
		Password: hash,
	}

	return s.repo.Create(&user)
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// check password
	if err := utils.CheckPassword(password, user.Password); err != nil {
		return "", errors.New("invalid credentials")
	}

	// generate JWT
	token, err := utils.GenerateToken(user.ID.String(), user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}