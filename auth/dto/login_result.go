package dto

import "learn_gqlgen/auth/entity"

type LoginResult struct {
	User    *entity.User
	Message string
}