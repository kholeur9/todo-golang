package services

import (
	"context"
	"learn_gqlgen/auth/dto"
	"learn_gqlgen/auth/entity"
	errors "learn_gqlgen/auth/errors_auth"
)

type AuthServicesImpl struct{}

func (asi *AuthServicesImpl) Create(ctx context.Context, input entity.User) (*dto.CreateUserResult, error) {
	if input.Name == "" && input.Email == "" && input.HashedPassword == "" {
		return nil, errors.ErrInfoNotCompleted
	}
	if len(input.Name) == 1 {
		return nil, errors.ErrNameNotTooLong
	}
	return &dto.CreateUserResult{}, nil
}