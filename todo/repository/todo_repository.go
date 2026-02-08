package repository

import (
	"learn_gqlgen/todo/entity"
)

type TodoRepository interface {
	Create(todo *entity.Todo) (*entity.Todo, error)
	FindTodoById(id string) (*entity.Todo, error)
	FindTodoByTextAndUserID(text string, userID string) (*entity.Todo, error)
	FindAllTodos(userID string) ([]*entity.Todo, error)
	Update(updateTodo *entity.UpdateTodo) (*entity.Todo, error)
	Delete(id string) error
}