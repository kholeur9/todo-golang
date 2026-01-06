package graph

import (
	todo "learn_gqlgen/todo/services"
	auth "learn_gqlgen/auth/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

type Resolver struct{
	TodoService todo.Services
	AuthService auth.Services
}
