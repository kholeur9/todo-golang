package services

import (
	"context"
	//"fmt"
	"learn_gqlgen/todo/entity"
	"learn_gqlgen/todo/errors"
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

func (tsi *TodoServiceImpl) AddTodo(ctx context.Context, input *entity.NewTodo) (*entity.Todo, error) {
	if input.Text == "" {
		return nil, errors.ErrTodoDonotEmpty
	}
	id := uuid.New().String()
	todoInitiated := &entity.Todo{
		ID: id,
		Text: input.Text,
		Done: false,
		User: &entity.User{ID: input.UserID, Name: "John Doe"},
	}
	todo, err := tsi.repo.Create(todoInitiated)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (tsi *TodoServiceImpl) GetTodo(ctx context.Context, id string) (*entity.Todo, error) {
	if id == "" {
		return nil, errors.ErrTodoIdEmpty
	}
	todoExists, err := tsi.repo.FindTodoById(id)
	if err != nil {
		return nil, err
	}
	return todoExists, nil
}

func (tsi *TodoServiceImpl) GetAllTodos(ctx context.Context) ([]*entity.Todo, error) {
	allTodos, err := tsi.repo.FindAllTodos()
	if err != nil {
		return nil, err
	}
	return allTodos, nil
}

func (tsi *TodoServiceImpl) UpdateTodo(ctx context.Context, input *entity.UpdateTodo) (*entity.Todo, error) {
	if input.Text == "" {
		return nil, errors.ErrTodoHasUpdateEmpty
	}
	id := input.TodoID
	todoExists, err := tsi.repo.FindTodoById(id)
	if err != nil {
		return nil, err
	}
	var todo *entity.Todo
	if todoExists.Text != input.Text && todoExists.Done == input.Done {
		if todoExists.Done {
			return nil, errors.ErrTodoAlreadyTrue
		}
		dataTextUpdate := &entity.UpdateTodo{
			TodoID: todoExists.ID,
			Text: input.Text,
			Done: todoExists.Done,
		}
		thisUpdateTodo, err := tsi.repo.Update(dataTextUpdate)
		if err != nil {
			return nil, err
		}
		todo = thisUpdateTodo
	} else if todoExists.Text == input.Text && todoExists.Done != input.Done {
		dataDoneUpdate := &entity.UpdateTodo{
			TodoID: todoExists.ID,
			Text: todoExists.Text,
			Done: input.Done,
		}
		thisUpdateTodo, err := tsi.repo.Update(dataDoneUpdate)
		if err != nil {
			return nil, err
		}
		todo = thisUpdateTodo
	} else if todoExists.Text != input.Text && todoExists.Done != input.Done {
		dataAllUpdate := &entity.UpdateTodo{
			TodoID: todoExists.ID,
			Text: input.Text,
			Done: false,
		}
		thisUpdate, err := tsi.repo.Update(dataAllUpdate)
		todo = thisUpdate
		if err != nil {
			return nil, err
		}
	} else {
		return todoExists, nil
	}
	return todo, nil
}