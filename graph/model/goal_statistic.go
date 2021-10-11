package model

import "time"

type GoalStatistic struct {
	ID           int        `json:"id"`
	GoalID       int        `json:"goalId"`
	Milliseconds int        `json:"milliseconds"`
	CreatedAt    *time.Time `json:"createdAt"`
	UpdatedAt    *time.Time `json:"updatedAt"`
}
