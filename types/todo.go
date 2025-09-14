package types

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeleteAt    gorm.DeletedAt `gorm:"index" json:"delete_at"`
	Done        bool           `json:"done"`
}

type CreateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTodoRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UpdatedAt   time.Time `json:"updated_at"`
	Done        bool      `json:"done"`
}

func NewTodo(tt, dc string) *Todo {
	return &Todo{
		Title:       tt,
		Description: dc,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Done:        false,
	}
}

func UpdateTodo(tt, dc string, dn bool) *Todo {
	return &Todo{
		Title:       tt,
		Description: dc,
		UpdatedAt:   time.Now().UTC(),
		Done:        dn,
	}
}
