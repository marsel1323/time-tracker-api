package model

import "time"

type Task struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Done       bool       `json:"done"`
	CategoryID int        `json:"categoryID"`
	CreatedAt  *time.Time `json:"createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt"`
}
