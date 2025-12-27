package memory

import (
	//"context"
	"fmt"
	"learn_gqlgen/todo/entity"
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
	return nil, fmt.Errorf("Todo non trouvé.")
}

func (ntmi *MemoryImpl) FindAllTodos() ([]*entity.Todo, error) {
	return ntmi.stockTodo, nil
}