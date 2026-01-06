package services

import (
	"context"
	"learn_gqlgen/auth/dto"
	"learn_gqlgen/auth/entity"
)

type Services interface {
	Create(ctx context.Context, input entity.User) (*dto.CreateUserResult, error)
}