package service

import (
	"github.com/marsel1323/timetrackerapi/repository"
)

type GoalStatisticService interface {
	TodayTime(goalId int) (int, error)
}

type goalStatisticService struct {
	repo repository.GoalStatisticRepository
}

func NewGoalStatisticService(repo repository.GoalStatisticRepository) *goalStatisticService {
	return &goalStatisticService{
		repo: repo,
	}
}

func (s *goalStatisticService) TodayTime(goalId int) (int, error) {
	return s.repo.TodayTime(goalId)
}
