package repository

import (
	"context"
	"database/sql"
	"github.com/marsel1323/timetrackerapi/graph/model"
	"time"
)

type ICategoryRepository interface {
	Create(category model.NewCategory) (*model.Category, error)
	List() ([]*model.Category, error)

	TotalTime(categoryId int) (int, error)
	TodayTime(categoryId int) (int, error)
}

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (repo *CategoryRepository) Create(newCategory model.NewCategory) (*model.Category, error) {
	query := `
		INSERT INTO categories (name)
		VALUES ($1)
		RETURNING id, name;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var category model.Category

	err := repo.db.QueryRowContext(ctx, query, newCategory.Name).Scan(
		&category.ID,
		&category.Name,
	)

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (repo *CategoryRepository) List() ([]*model.Category, error) {
	query := `
		SELECT id, name
		FROM categories
		ORDER BY id;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*model.Category

	for rows.Next() {
		var category model.Category

		err := rows.Scan(
			&category.ID,
			&category.Name,
		)
		if err != nil {
			return nil, err
		}

		categories = append(categories, &category)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func (repo *CategoryRepository) TotalTime(categoryId int) (int, error) {
	query := `
		SELECT coalesce(sum(s.milliseconds), 0)
		FROM task LEFT JOIN stats s on task.id = s.task_id
		WHERE category_id = $1
		GROUP BY task.category_id;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var sum = 0

	err := repo.db.QueryRowContext(ctx, query, categoryId).Scan(&sum)
	if err != nil {
		return 0, err
	}

	return sum, nil
}

func (repo *CategoryRepository) TodayTime(categoryId int) (int, error) {
	query := `
		SELECT coalesce(sum(s.milliseconds), 0)
		FROM task LEFT JOIN stats s on task.id = s.task_id
		WHERE category_id = $1 AND s.created_at::date = $2
		GROUP BY task.category_id;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var sum = 0
	currentTime := time.Now().Format("2006-01-02")

	err := repo.db.QueryRowContext(ctx, query, categoryId, currentTime).Scan(&sum)
	if err != nil {
		return 0, err
	}

	return sum, nil
}
