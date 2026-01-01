package memory

import (
	//"context"
	//"fmt"
	"learn_gqlgen/todo/entity"
	"learn_gqlgen/todo/errors"
	//"learn_gqlgen/todo/repository"
	//"learn_gqlgen/todo/repository"
)


type MemoryImpl struct{
	stockTodo []*entity.Todo
}

func NewTodoRepository(stockTodo []*entity.Todo) *MemoryImpl {
	return &MemoryImpl{
		stockTodo: stockTodo,
	}
}

func (ntmi *MemoryImpl) Create(todo *entity.Todo) (*entity.Todo, error) {
	ntmi.stockTodo = append(ntmi.stockTodo, todo)
	return todo, nil
}

func (ntmi *MemoryImpl) FindTodoById(id string) (*entity.Todo, error) {
	for _, thisTodo := range ntmi.stockTodo {
		if  thisTodo.ID == id {
			return thisTodo, nil
		}
	}
	return nil, errors.ErrTodoNotFound
}

func (ntmi *MemoryImpl) FindAllTodos() ([]*entity.Todo, error) {
	return ntmi.stockTodo, nil
}

func (ntmi *MemoryImpl) Update(updateTodo *entity.UpdateTodo) (*entity.Todo, error) {
	for _, todo := range ntmi.stockTodo {
		if todo.ID == updateTodo.TodoID {
			todo.Text = updateTodo.Text
			todo.Done = updateTodo.Done
			return todo, nil
		}
	}
	return nil, errors.ErrTodoNoUpdate
}

func (ntmi *MemoryImpl) Delete(id string) error {
	for i, todo := range ntmi.stockTodo {
		if todo.ID == id {
			ntmi.stockTodo = append(ntmi.stockTodo[:i], ntmi.stockTodo[i+1:]...)
			return nil
		}
	}
	return errors.ErrTodoDoNotDeleted
}