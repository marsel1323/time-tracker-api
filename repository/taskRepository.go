package repository

import (
	"context"
	"database/sql"
	"github.com/marsel1323/timetrackerapi/graph/model"
	"time"
)

type TaskRepository interface {
	Create(task *model.NewTask) (*model.Task, error)
	List() ([]*model.Task, error)
	ListByCategory(categoryId int) ([]*model.Task, error)
	Get(id int) (*model.Task, error)
}

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *taskRepository {
	return &taskRepository{
		db: db,
	}
}

func (repo *taskRepository) Create(newTask *model.NewTask) (*model.Task, error) {
	query := `
		INSERT INTO task (name, category_id)
		VALUES ($1, $2)
		RETURNING id, name, category_id, created_at, updated_at;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var task model.Task

	err := repo.db.QueryRowContext(ctx, query, newTask.Name).Scan(
		&task.ID,
		&task.Name,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (repo *taskRepository) List() ([]*model.Task, error) {
	query := `
		SELECT id, name, done, created_at, updated_at
		FROM task
		WHERE done = false
		ORDER BY id;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*model.Task

	for rows.Next() {
		var task model.Task

		err := rows.Scan(
			&task.ID,
			&task.Name,
			&task.Done,
			&task.CreatedAt,
			&task.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, &task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (repo *taskRepository) ListByCategory(categoryId int) ([]*model.Task, error) {
	query := `
		SELECT id, name, category_id, created_at, updated_at
		FROM task
		WHERE category_id = $1
		ORDER BY id;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := repo.db.QueryContext(ctx, query, categoryId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*model.Task

	for rows.Next() {
		var task model.Task

		err := rows.Scan(
			&task.ID,
			&task.Name,
			&task.CategoryID,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, &task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (repo *taskRepository) Get(id int) (*model.Task, error) {
	query := `
		SELECT id, name, created_at, updated_at
		FROM task
		WHERE id = $1;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var task model.Task

	err := repo.db.QueryRowContext(ctx, query, id).Scan(
		&task.ID,
		&task.Name,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &task, nil
}
