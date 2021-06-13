package model

import "time"

type Task struct {
	ID               int        `json:"id"`
	Name             string     `json:"name"`
	CategoryID       int        `json:"categoryId"`
	CreatedAt        *time.Time `json:"createdAt"`
	UpdatedAt        *time.Time `json:"updatedAt"`
}
