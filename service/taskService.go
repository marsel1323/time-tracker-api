package service

import (
	"github.com/marsel1323/timetrackerapi/graph/model"
	"github.com/marsel1323/timetrackerapi/repository"
)

type TaskService interface {
	CreateTask(input *model.NewTask) (*model.Task, error)
	TaskList() ([]*model.Task, error)
	TaskListByCategory(categoryId int) ([]*model.Task, error)
	GetTask(id int) (*model.Task, error)
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *taskService {
	return &taskService{
		repo: repo,
	}
}

func (s *taskService) CreateTask(input *model.NewTask) (*model.Task, error) {
	return s.repo.Create(input)
}

func (s *taskService) TaskList() ([]*model.Task, error) {
	return s.repo.List()
}

func (s *taskService) TaskListByCategory(categoryId int) ([]*model.Task, error) {
	return s.repo.ListByCategory(categoryId)
}

func (s *taskService) GetTask(id int) (*model.Task, error) {
	return s.repo.Get(id)
}
