package services

import (
	"context"
	"crypto/rand"
	"fmt"
	"learn_gqlgen/todo/entity"
	"math/big"
)

type TodoService interface {
	AddTodo(ctx context.Context, input entity.NewTodo) (*entity.Todo, error)
}

type TodoServiceImpl struct {}

func NewTodoService() *TodoServiceImpl {
	return &TodoServiceImpl{}
}

func (ti TodoServiceImpl) AddTodo(ctx context.Context, input entity.NewTodo) (*entity.Todo, error) {
	id, _ := rand.Int(rand.Reader, big.NewInt(100))
	return &entity.Todo{
		ID: fmt.Sprintf("T&d", id),
		Text: input.Text,
		Done: false,
		User: &entity.User{ID: input.UserID, Name: "John Doe"},
	}, nil
}