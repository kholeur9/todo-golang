package repository

import (
	"learn_gqlgen/todo/entity"
)

type TodoRepository interface {
	Create(todo *entity.Todo) (*entity.Todo, error)
	FindTodoById(id string) (*entity.Todo, error)
	FindAllTodos() ([]*entity.Todo, error)
}