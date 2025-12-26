package services

import (
	"context"
	"learn_gqlgen/todo/entity"
)

type Services interface {
	AddTodo(ctx context.Context, input entity.NewTodo) (*entity.Todo, error)
}