package service

import (
	"github.com/marsel1323/timetrackerapi/graph/model"
	"github.com/marsel1323/timetrackerapi/repository"
)

type GoalService interface {
	CreateGoal(input model.NewGoal) (*model.Goal, error)
	GoalList() ([]*model.Goal, error)
}

type goalService struct {
	repo repository.GoalRepository
}

func NewGoalService(repo repository.GoalRepository) *goalService {
	return &goalService{
		repo: repo,
	}
}

func (s *goalService) CreateGoal(input model.NewGoal) (*model.Goal, error) {
	return s.repo.Create(input)
}

func (s *goalService) GoalList() ([]*model.Goal, error) {
	return s.repo.List()
}
