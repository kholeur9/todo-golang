package services

import (
	"context"
	"crypto/rand"
	"fmt"
	"learn_gqlgen/todo/entity"
	"learn_gqlgen/todo/repository"
	//"learn_gqlgen/todo/repository/memory"
	"math/big"
)

type TodoService interface {
	AddTodo(ctx context.Context, input entity.NewTodo) (*entity.Todo, error)
}

type TodoServiceImpl struct {
	inMemory repository.TodoRepository
}

func NewTodoService(inMemory repository.TodoRepository) *TodoServiceImpl {
	return &TodoServiceImpl{inMemory: inMemory}
}

func (ti *TodoServiceImpl) AddTodo(ctx context.Context, input entity.NewTodo) (*entity.Todo, error) {
	id, _ := rand.Int(rand.Reader, big.NewInt(100))
	todo := &entity.Todo{
		ID: fmt.Sprintf("T&d", id),
		Text: input.Text,
		Done: false,
		User: &entity.User{ID: input.UserID, Name: "John Doe"},
	}
	err := ti.inMemory.Create(todo)

	if err != nil {
		return nil, fmt.Errorf("Erreur lors de l'insertion de la todo.")
	}
	return todo, nil
}