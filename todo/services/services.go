package services

import (
	"context"
	"learn_gqlgen/todo/entity"
)

type Services interface {
	AddTodo(ctx context.Context, input *entity.NewTodo) (*entity.Todo, error)
	GetTodo(ctx context.Context, id string) (*entity.Todo, error)
	GetAllTodos(ctx context.Context) ([]*entity.Todo, error)
}