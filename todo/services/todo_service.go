package services

import (
	"context"
	"errors"
	"fmt"
	"strings"

	//"strconv"
	"time"

	//"fmt"
	"learn_gqlgen/todo/dto"
	"learn_gqlgen/todo/entity"
	"learn_gqlgen/todo/errors_todo"
	"learn_gqlgen/todo/repository"

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
		return nil, errors_todo.ErrTodoDonotEmpty
	}
	if len(input.Text) > 200 {
		return nil, errors_todo.ErrTodoTextTooLong
	}
	text := strings.TrimSpace(input.Text)
	sameTodoText, err := tsi.repo.FindTodoByText(text)
	if err != nil {
		if !errors.Is(err, errors_todo.ErrTodoNotFound) {
			return nil, err
		}
	} else {
		if sameTodoText != nil {
			return nil, errors_todo.ErrTodoDuplicate
		}
	}

	id := uuid.New().String()
	todoInitiated := &entity.Todo{
		ID: id,
		Text: text,
		Done: false,
		User: &entity.User{ID: input.UserID, Name: "John Doe"},
		CreatedAt: time.Now(),
		UpdatedAt: nil,
	}
	todo, err := tsi.repo.Create(todoInitiated)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (tsi *TodoServiceImpl) GetTodo(ctx context.Context, id string) (*entity.Todo, error) {
	if id == "" {
		return nil, errors_todo.ErrTodoIdEmpty
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

func (tsi *TodoServiceImpl) UpdateTodo(ctx context.Context, input *entity.UpdateTodo) (*dto.TodoUpdateResult, error) {
	if input.Text == "" {
		return nil, errors_todo.ErrTodoHasUpdateEmpty
	}
	id := input.TodoID
	todoExists, err := tsi.repo.FindTodoById(id)
	if err != nil {
		return nil, err
	}
	var todo *entity.Todo
	var message string
	now := time.Now()
	if todoExists.Text != input.Text && todoExists.Done == input.Done {
		if todoExists.Done {
			return nil, errors_todo.ErrTodoAlreadyTrue
		}
		dataTextUpdate := &entity.UpdateTodo{
			TodoID: todoExists.ID,
			Text: input.Text,
			Done: todoExists.Done,
			UpdatedAt: &now,
		}
		thisUpdateTodo, err := tsi.repo.Update(dataTextUpdate)
		if err != nil {
			return nil, err
		}
		todo = thisUpdateTodo
		message = "Votre todo a été modifié avec succès."
	} else if todoExists.Text == input.Text && todoExists.Done != input.Done {
		dataDoneUpdate := &entity.UpdateTodo{
			TodoID: todoExists.ID,
			Text: todoExists.Text,
			Done: input.Done,
			UpdatedAt: &now,
		}
		thisUpdateTodo, err := tsi.repo.Update(dataDoneUpdate)
		if err != nil {
			return nil, err
		}
		todo = thisUpdateTodo
		message = "Votre todo a été modifié avec succès."
	} else if todoExists.Text != input.Text && todoExists.Done != input.Done {
		dataAllUpdate := &entity.UpdateTodo{
			TodoID: todoExists.ID,
			Text: input.Text,
			Done: false,
			UpdatedAt: &now,
		}
		thisUpdate, err := tsi.repo.Update(dataAllUpdate)
		if err != nil {
			return nil, err
		}
		todo = thisUpdate
		message = "La tâche a été remise à false suite à la modification de son contenu."
	} else {
		return &dto.TodoUpdateResult{
			Todo: todoExists,
			Message: "Aucune modification n'a été appliqué. Vous n'avez rien modifié.",
		}, nil
	}
	return &dto.TodoUpdateResult{
		Todo: todo,
		Message: message,
	}, nil
}

func (tsi *TodoServiceImpl) DeleteTodo(ctx context.Context, id string) (*dto.TodoDeleteResult, error) {
	if id == "" {
		return nil, errors_todo.ErrTodoIdEmpty
	}
	getThisTodo, err := tsi.repo.FindTodoById(id)
	if err != nil {
		return nil, err
	}
	err = tsi.repo.Delete(getThisTodo.ID)
	if err != nil {
		return nil, err
	}
	return &dto.TodoDeleteResult{
		ID: id,
		Message: "Todo supprimée avec succès.",
	}, nil
}

func (tsi *TodoServiceImpl) DeleteTodos(ctx context.Context, IDs []string) (*dto.TodosDeleteResult, error) {
	fmt.Println("J'ai reçu les ids à supprimer", IDs)
	idsDeleted := []string{}
	idsNotDeleted := []string{}

	if len(IDs) == 0 {
		return nil, errors_todo.ErrTodoIdEmpty
	}
	for _, id := range IDs {
		err := tsi.repo.Delete(id)
		if err == nil {
			idsDeleted = append(idsDeleted, id)
		}
		if err != nil {
			idsNotDeleted = append(idsNotDeleted, id)
		}
	}
	message := fmt.Sprintf("%v supprimées, %v non supprimées.", len(idsDeleted), len(idsNotDeleted))
	return &dto.TodosDeleteResult{
		IDsDeleted: idsDeleted,
		IDsNotDeleted: idsNotDeleted,
		Message: message,
	}, nil
}