package memory

import (
	//"context"
	"fmt"
	"learn_gqlgen/todo/entity"
	//"learn_gqlgen/todo/repository"
	//"learn_gqlgen/todo/repository"
)


type MemoryImpl struct{
	stockTodo map[string]*entity.Todo
}

func NewTodoRepository() *MemoryImpl {
	return &MemoryImpl{}
}

func (ntmi *MemoryImpl) Create(todo *entity.Todo) error {
	fmt.Println("J'ai reçu la todo")
	return nil
}