package dto

import "learn_gqlgen/todo/entity"

type TodoUpdateResult struct {
	Todo *entity.Todo
	Message string
}

type TodoDeleteResult struct {
	Message string
}