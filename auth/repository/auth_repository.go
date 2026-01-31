package repository

import "learn_gqlgen/auth/entity"

type UserRepository interface {
	CreateAccount(user *entity.User) (*entity.User, error)
	FindUserByEmail(email string) (*entity.User, error)
}