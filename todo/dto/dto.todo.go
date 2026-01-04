package dto

import "learn_gqlgen/todo/entity"

type TodoUpdateResult struct {
	Todo *entity.Todo
	Message string
}

type TodoDeleteResult struct {
	ID string
	Message string
}

type TodosDeleteResult struct {
	IDsDeleted []string
	IDsNotDeleted []string
	Message string
}