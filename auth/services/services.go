package services

import (
	"context"
	"learn_gqlgen/auth/dto"
	//"learn_gqlgen/auth/entity"
)

type Services interface {
	Create(ctx context.Context, input dto.RegisterInput) (*dto.CreateUserResult, error)
	Login(ctx context.Context, input dto.LoginInput) (*dto.LoginResult, error)
}