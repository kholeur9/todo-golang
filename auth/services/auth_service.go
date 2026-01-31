package services

import (
	"context"
	"errors"
	//"fmt"

	//"learn_gqlgen/auth/dto"
	"learn_gqlgen/auth/dto"
	"learn_gqlgen/auth/entity"
	errors_auth "learn_gqlgen/auth/errors_auth"
	"learn_gqlgen/auth/repository"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthServicesImpl struct{
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) *AuthServicesImpl {
	return &AuthServicesImpl{
		repo: repo,
	}
}

func (asi *AuthServicesImpl) Create(ctx context.Context, input dto.RegisterInput) (*dto.CreateUserResult, error) {
	if input.Name == "" || input.Email == "" || input.Password == "" {
		return nil, errors_auth.ErrInfoNotCompleted
	}
	if len(input.Name) == 1 {
		return nil, errors_auth.ErrNameNotTooLong
	}
	emailExists, err := asi.repo.FindUserByEmail(input.Email)
	if err != nil {
		if !errors.Is(err, errors_auth.ErrIncorrectCredentials) {
			return nil, err
		}
	} else {
		if emailExists != nil {
			return nil, errors_auth.ErrEmailAlreadyExists
		}
	}
	if len(input.Password) < 8 {
		return nil, errors_auth.ErrPasswordTooShort
	}
	// Create User
	id := uuid.New().String()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	initiatedUser := &entity.User{
		ID: id,
		Name: input.Name,
		Email: input.Email,
		HashedPassword: string(hashedPassword),
		CreatedAt: time.Now(),
	}
	message := "Votre compte a été créée avec succès."
	user, err := asi.repo.CreateAccount(initiatedUser)
	if err != nil {
		return nil, err
	}
	return &dto.CreateUserResult{
		User: user,
		Message: message,
	}, nil
}

func (asi *AuthServicesImpl) Login(ctx context.Context, input dto.LoginInput) (*dto.LoginResult, error) {
	if input.Email == "" || input.Password == "" {
		return nil, errors_auth.ErrInfoNotCompleted
	}
	user, err := asi.repo.FindUserByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(input.Password))
	if err != nil {
		return nil, errors_auth.ErrIncorrectCredentials
	}
	// Login User
	return &dto.LoginResult{
		User: &entity.User{
			ID: user.ID,
			Name: user.Name,
			Email: user.Email,
		},
		Message: "",
	}, nil
}