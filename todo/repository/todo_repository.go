package repository

import (
	"learn_gqlgen/todo/entity"
)

type TodoRepository interface {
	Create(todo *entity.Todo) (*entity.Todo, error)
	FindTodoById(id string) (*entity.Todo, error)
	FindTodoByText(text string) (*entity.Todo, error)
	FindAllTodos() ([]*entity.Todo, error)
	Update(updateTodo *entity.UpdateTodo) (*entity.Todo, error)
	Delete(id string) error
}