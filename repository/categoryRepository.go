package repository

import (
	"context"
	"database/sql"
	"github.com/marsel1323/timetrackerapi/graph/model"
	"time"
)

type CategoryRepository interface {
	Create(category model.NewCategory) (*model.Category, error)
	List() ([]*model.Category, error)

	TotalTime(categoryId int) (int, error)
	TodayTime(categoryId int) (int, error)
	TimeByDate(categoryId int, date *string) (int, error)
}

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *categoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (repo *categoryRepository) Create(newCategory model.NewCategory) (*model.Category, error) {
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

func (repo *categoryRepository) List() ([]*model.Category, error) {
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

func (repo *categoryRepository) TotalTime(categoryId int) (int, error) {
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

func (repo *categoryRepository) TodayTime(categoryId int) (int, error) {
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

func (repo *categoryRepository) TimeByDate(categoryId int, date *string) (int, error) {
	query := `
		SELECT coalesce(sum(s.milliseconds), 0)
		FROM task LEFT JOIN stats s on task.id = s.task_id
		WHERE category_id = $1 AND s.created_at::date = $2
		GROUP BY task.category_id;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var sum = 0

	var today string

	if date == nil {
		today = time.Now().Format("2006-01-02")
	} else {
		today = *date
	}

	parsedDate, err := time.Parse("2006-01-02", today)
	if err != nil {
		//parsedDate = time.Now()//.Format("2006-01-02")
		return 0, err
	}

	err = repo.db.QueryRowContext(ctx, query, categoryId, parsedDate).Scan(&sum)
	if err != nil {
		return 0, err
	}

	return sum, nil
}
