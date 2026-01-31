package services

import (
	"context"
	"learn_gqlgen/auth/dto"
	"learn_gqlgen/auth/entity"
)

type Services interface {
	Create(ctx context.Context, input *entity.NewUser) (*dto.CreateUserResult, error)
}