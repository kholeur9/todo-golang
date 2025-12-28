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

func (ntmi *MemoryImpl) Update(updateTodo *entity.UpdateTodo) (*entity.Todo, error) {
	fmt.Println("J'ai reçu la todo à mettre à jour.")
	for _, todo := range ntmi.stockTodo {
		if updateTodo.TodoID == todo.ID {
			todo.Text = updateTodo.Text
			todo.Done = updateTodo.Done
			return todo, nil
		}
	}
	return nil, fmt.Errorf("Todo non mise à jour.")
}