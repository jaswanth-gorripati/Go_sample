package model

import "time"

type Todo struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    int       `json:"user_id" gorm:"index"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TodoInput struct {
	Title string `json:"title"`
	Done  *bool  `json:"done"`
}
