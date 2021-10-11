package service

import (
	"github.com/marsel1323/timetrackerapi/graph/model"
	"github.com/marsel1323/timetrackerapi/repository"
	"time"
)

type IStatisticService interface {
	CreateStat(input model.NewTaskStatistic) (*model.TaskStatistic, error)
	StatList() ([]*model.TaskStatistic, error)
	StatListByDate(date string) ([]*model.TaskStatistic, error)
	GetStat(id int) (*model.TaskStatistic, error)

	CalcTotalTime(taskId int) (*int, error)
	CalcTotalTodayTime(taskID int) (*int, error)
	CalcTotalTimeFor(taskId int, days int, hours int) (int, error)
	LastStatRecord(taskId int) (*model.TaskStatistic, error)
	TotalTimeFor(taskId int, day string) (int, error)
}

type StatisticService struct {
	repo repository.StatisticRepository
}

func NewStatisticService(repo repository.StatisticRepository) *StatisticService {
	return &StatisticService{
		repo: repo,
	}
}

func (s *StatisticService) CreateStat(input model.NewTaskStatistic) (*model.TaskStatistic, error) {
	newStat := model.TaskStatistic{
		Milliseconds: input.Ms,
		TaskID:       input.TaskID,
	}

	return s.repo.Create(newStat)
}

func (s *StatisticService) StatList() ([]*model.TaskStatistic, error) {
	return s.repo.List()
}

func (s *StatisticService) StatListByDate(date string) ([]*model.TaskStatistic, error) {
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}
	return s.repo.ListByDate(date)
}

func (s *StatisticService) GetStat(id int) (*model.TaskStatistic, error) {
	return s.repo.Get(id)
}

func (s *StatisticService) CalcTotalTime(taskId int) (*int, error) {
	return s.repo.TaskTotalTime(taskId)
}

func (s *StatisticService) CalcTotalTodayTime(taskId int) (*int, error) {
	return s.repo.TaskTodayTime(taskId)
}

func (s *StatisticService) CalcTotalTimeFor(taskId int, days int, hours int) (int, error) {
	return s.repo.TaskTotalTimeFor(taskId, days, hours)
}

func (s *StatisticService) LastStatRecord(taskId int) (*model.TaskStatistic, error) {
	return s.repo.LastRecord(taskId)
}

func (s *StatisticService) TotalTimeFor(taskId int, day string) (int, error) {
	return s.repo.TaskTotalTimeForDay(taskId, day)
}
