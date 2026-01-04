package memory

import (
	//"fmt"
	//"slices"
	//"fmt"
	"learn_gqlgen/todo/entity"
	"learn_gqlgen/todo/errors_todo"
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
	return nil, errors_todo.ErrTodoNotFound
}

func (ntmi *MemoryImpl) FindTodoByText(text string) (*entity.Todo, error) {
	for _, todo := range ntmi.stockTodo {
		if todo.Text == text {
			return todo, nil
		}
	}
	return nil, errors_todo.ErrTodoNotFound
}

func (ntmi *MemoryImpl) FindAllTodos() ([]*entity.Todo, error) {
	return ntmi.stockTodo, nil
}

func (ntmi *MemoryImpl) Update(updateTodo *entity.UpdateTodo) (*entity.Todo, error) {
	for _, todo := range ntmi.stockTodo {
		if todo.ID == updateTodo.TodoID {
			todo.Text = updateTodo.Text
			todo.Done = updateTodo.Done
			todo.UpdatedAt = updateTodo.UpdatedAt
			return todo, nil
		}
	}
	return nil, errors_todo.ErrTodoNoUpdate
}

func (ntmi *MemoryImpl) Delete(id string) error {
	for i, todo := range ntmi.stockTodo {
		if todo.ID == id {
			ntmi.stockTodo = append(ntmi.stockTodo[:i], ntmi.stockTodo[i+1:]...)
			return nil
		}
	}
	return errors_todo.ErrTodoDoNotDeleted
}
