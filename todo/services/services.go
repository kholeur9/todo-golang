package services

import (
	"context"
	"learn_gqlgen/todo/dto"
	"learn_gqlgen/todo/entity"
)

type Services interface {
	AddTodo(ctx context.Context, input *entity.NewTodo) (*entity.Todo, error)
	GetTodo(ctx context.Context, id string) (*entity.Todo, error)
	GetAllTodos(ctx context.Context, userID string) ([]*entity.Todo, error)
	UpdateTodo(ctx context.Context, input *entity.UpdateTodo) (*dto.TodoUpdateResult, error)
	DeleteTodo(ctx context.Context, id string) (*dto.TodoDeleteResult, error)
	DeleteTodos(cts context.Context, IDs []string) (*dto.TodosDeleteResult, error)
}