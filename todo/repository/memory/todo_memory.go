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
	fmt.Println("J'ai reçu la todo")
	ntmi.stockTodo = append(ntmi.stockTodo, todo)
	fmt.Println("Memory : ", todo)
	return todo, nil
}