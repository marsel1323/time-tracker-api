package service

import (
	"github.com/marsel1323/timetrackerapi/graph/model"
	"github.com/marsel1323/timetrackerapi/repository"
)

type ICategoryService interface {
	CreateCategory(newCategory model.NewCategory) (*model.Category, error)
	List() ([]*model.Category, error)

	TotalTime(categoryId int) (int, error)
	TodayTime(categoryId int) (int, error)
}

type CategoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (s *CategoryService) CreateCategory(newCategory model.NewCategory) (*model.Category, error) {
	return s.repo.Create(newCategory)
}

func (s *CategoryService) List() ([]*model.Category, error) {
	return s.repo.List()
}

func (s *CategoryService) TotalTime(categoryId int) (int, error) {
	return s.repo.TotalTime(categoryId)
}

func (s *CategoryService) TodayTime(categoryId int) (int, error) {
	return s.repo.TodayTime(categoryId)
}
