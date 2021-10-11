// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewCategory struct {
	Name string `json:"name"`
}

type NewGoal struct {
	Name       string `json:"name"`
	CategoryID int    `json:"categoryId"`
	Time       int    `json:"time"`
}

type NewTask struct {
	Name string `json:"name"`
}

type NewTaskStatistic struct {
	TaskID int `json:"taskId"`
	Ms     int `json:"ms"`
}
