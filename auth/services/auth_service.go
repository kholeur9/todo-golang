package services

import (
	"context"
	"fmt"

	//"learn_gqlgen/auth/dto"
	"learn_gqlgen/auth/dto"
	"learn_gqlgen/auth/entity"
	errors "learn_gqlgen/auth/errors_auth"
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

func (asi *AuthServicesImpl) Create(ctx context.Context, input entity.NewUser) (*dto.CreateUserResult, error) {
	if input.Name == "" && input.Email == "" && input.Password == "" {
		return nil, errors.ErrInfoNotCompleted
	}
	if len(input.Name) == 1 {
		return nil, errors.ErrNameNotTooLong
	}
	emailExists, err := asi.repo.FindUserByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	if emailExists != nil {
		return nil, errors.ErrEmailAlreadyExists
	}
	if len(input.Password) < 8 {
		return nil, errors.ErrPasswordTooShort
	}
	// Create User
	id := uuid.New().String()
	hashedPassword, _ := (bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost))
	fmt.Println("Hashed Paswword: ", string(hashedPassword))
	initiatedUser := &entity.User{
		ID: id,
		Name: input.Name,
		Email: input.Email,
		HashedPassword: string(hashedPassword),
		CreatedAt: time.Now(),
	}
	user, err := asi.repo.CreateAccount(initiatedUser)
	if err != nil {
		return nil, err
	}
	return &dto.CreateUserResult{
		User: user,
		Message: "",
	}, nil
}