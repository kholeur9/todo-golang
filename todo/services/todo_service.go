package services

import (
	"context"
	"fmt"
	"learn_gqlgen/todo/entity"
	"learn_gqlgen/todo/repository"

	//"learn_gqlgen/todo/repository/memory"

	"github.com/google/uuid"
)

type TodoServiceImpl struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) *TodoServiceImpl {
	return &TodoServiceImpl{
		repo: repo,
	}
}

func (tsi *TodoServiceImpl) AddTodo(ctx context.Context, input entity.NewTodo) (*entity.Todo, error) {
	id := uuid.New().String()
	todoInitiated := &entity.Todo{
		ID: id,
		Text: input.Text,
		Done: false,
		User: &entity.User{ID: input.UserID, Name: "John Doe"},
	}
	todo, err := tsi.repo.Create(todoInitiated)
	if err != nil {
		return nil, fmt.Errorf("Erreur lors de l'insertion de la todo.")
	}
	return todo, nil
}