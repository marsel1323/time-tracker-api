package model

import "time"

type Goal struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Time       int        `json:"time"`
	CategoryID int        `json:"categoryId"`
	CreatedAt  *time.Time `json:"createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt"`
}
