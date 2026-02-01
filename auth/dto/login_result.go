package dto

import "learn_gqlgen/auth/entity"

type LoginResult struct {
	User    *entity.User
	AccessToken string
	Message string
}