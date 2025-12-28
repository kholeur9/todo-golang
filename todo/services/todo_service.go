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

func (tsi *TodoServiceImpl) AddTodo(ctx context.Context, input *entity.NewTodo) (*entity.Todo, error) {
	if input.Text == "" {
		return nil, fmt.Errorf("La todo ne peut pas étre vide.")
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
		return nil, fmt.Errorf("Erreur lors de l'insertion de la todo.")
	}
	return todo, nil
}

func (tsi *TodoServiceImpl) GetTodo(ctx context.Context, id string) (*entity.Todo, error) {
	todoExists, err := tsi.repo.FindTodoById(id)
	if err != nil {
		return nil, fmt.Errorf("Erreur lors de la récupération de la Todo.")
	}
	return todoExists, nil
}

func (tsi *TodoServiceImpl) GetAllTodos(ctx context.Context) ([]*entity.Todo, error) {
	allTodos, err := tsi.repo.FindAllTodos()
	if err != nil {
		return nil, fmt.Errorf("Impossible de récupérer les todos.")
	}
	return allTodos, nil
}

func (tsi *TodoServiceImpl) UpdateTodo(ctx context.Context, input *entity.UpdateTodo) (*entity.Todo, error) {
	if input.Text == "" {
		return nil, fmt.Errorf("Veuillez indiquer une tâche à faire.")
	}
	id := input.TodoID
	todoExists, err := tsi.repo.FindTodoById(id)
	if err != nil {
		return nil, fmt.Errorf("Impossoble de mettre à jour la Todo.")
	}
	var todo *entity.Todo
	if todoExists.Text != input.Text && todoExists.Done == input.Done {
		dataTextUpdate := &entity.UpdateTodo{
			TodoID: todoExists.ID,
			Text: input.Text,
			Done: todoExists.Done,
		}
		thisUpdateTodo, err := tsi.repo.Update(dataTextUpdate)
		if err != nil {
			return nil, fmt.Errorf("Impossible de mettre à jour la todo.")
		}
		todo = thisUpdateTodo
		fmt.Println("Service : ", todo)
	} else if todoExists.Text == input.Text && todoExists.Done != input.Done {
		dataDoneUpdate := &entity.UpdateTodo{
			TodoID: todoExists.ID,
			Text: todoExists.Text,
			Done: input.Done,
		}
		thisUpdateTodo, err := tsi.repo.Update(dataDoneUpdate)
		if err != nil {
			return nil, fmt.Errorf("Impossible de mettre à jour la todo.")
		}
		todo = thisUpdateTodo
	} else {
		return todoExists, nil
	}
	return todo, nil
}