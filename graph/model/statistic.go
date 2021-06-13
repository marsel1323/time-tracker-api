package model

import "time"

type Statistic struct {
	ID           int        `json:"id"`
	TaskID       int        `json:"taskId"`
	Milliseconds int        `json:"milliseconds"`
	CreatedAt    *time.Time `json:"createdAt"`
	UpdatedAt    *time.Time `json:"updatedAt"`
}