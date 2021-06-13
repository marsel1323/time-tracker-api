package service

import (
	"github.com/marsel1323/timetrackerapi/graph/model"
	"github.com/marsel1323/timetrackerapi/repository"
)

type ITaskService interface {
	CreateTask(input *model.NewTask) (*model.Task, error)
	TaskList() ([]*model.Task, error)
	TaskListByCategory(categoryId int) ([]*model.Task, error)
	GetTask(id int) (*model.Task, error)
}

type TaskService struct {
	repo repository.ITaskRepository
}

func NewTaskService(repo repository.ITaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (s *TaskService) CreateTask(input *model.NewTask) (*model.Task, error) {
	return s.repo.Create(input)
}

func (s *TaskService) TaskList() ([]*model.Task, error) {
	return s.repo.List()
}

func (s *TaskService) TaskListByCategory(categoryId int) ([]*model.Task, error) {
	return s.repo.ListByCategory(categoryId)
}

func (s *TaskService) GetTask(id int) (*model.Task, error) {
	return s.repo.Get(id)
}
