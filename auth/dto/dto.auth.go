package dto

import (
	"learn_gqlgen/auth/entity"
)

type CreateUserResult struct {
	User *entity.User
	Message string
}