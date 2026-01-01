package entity

import "time"

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type UpdateTodo struct {
	TodoID string
	Text   string
	Done   bool
	UpdatedAt *time.Time
}

type Todo struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	Done      bool      `json:"done"`
	User      *User     `json:"user"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}