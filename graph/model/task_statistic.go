package model

import "time"

type TaskStatistic struct {
	ID           int        `json:"id"`
	TaskID       int        `json:"taskId"`
	Milliseconds int        `json:"milliseconds"`
	CreatedAt    *time.Time `json:"createdAt"`
	UpdatedAt    *time.Time `json:"updatedAt"`
}