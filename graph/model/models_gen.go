// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewCategory struct {
	Name string `json:"name"`
}

type NewStatistic struct {
	Ms     int `json:"ms"`
	TaskID int `json:"taskId"`
}

type NewTask struct {
	Name       string `json:"name"`
	CategoryID int    `json:"categoryId"`
}