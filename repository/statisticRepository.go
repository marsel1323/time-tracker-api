package repository

import (
	"context"
	"database/sql"
	"github.com/marsel1323/timetrackerapi/graph/model"
	"time"
)

type StatisticRepository interface {
	Create(stat model.TaskStatistic) (*model.TaskStatistic, error)
	List() ([]*model.TaskStatistic, error)
	ListByDate(date string) ([]*model.TaskStatistic, error)
	Get(id int) (*model.TaskStatistic, error)
	Update(stat *model.TaskStatistic) (*model.TaskStatistic, error)

	TaskTotalTime(taskId int) (*int, error)
	TaskTodayTime(taskId int) (*int, error)
	TaskTotalTimeFor(taskId int, days int, hours int) (int, error)
	TaskTotalTimeForDay(taskId int, day string) (int, error)
	LastRecord(taskId int) (*model.TaskStatistic, error)
}

type statisticRepository struct {
	db *sql.DB
}

func NewStatisticRepository(db *sql.DB) *statisticRepository {
	return &statisticRepository{
		db: db,
	}
}

func (repo *statisticRepository) Create(stat model.TaskStatistic) (*model.TaskStatistic, error) {
	query := `
		INSERT INTO stats (milliseconds, task_id)
		VALUES ($1, $2)
		RETURNING id, created_at, updated_at;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []interface{}{stat.Milliseconds, stat.TaskID}

	err := repo.db.QueryRowContext(ctx, query, args...).Scan(
		&stat.ID,
		&stat.CreatedAt,
		&stat.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &stat, nil
}

func (repo *statisticRepository) List() ([]*model.TaskStatistic, error) {
	query := `
		SELECT id, milliseconds, task_id, created_at, updated_at
		FROM stats;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats []*model.TaskStatistic

	for rows.Next() {
		var stat model.TaskStatistic

		err := rows.Scan(
			&stat.ID,
			&stat.Milliseconds,
			&stat.TaskID,
			&stat.CreatedAt,
			&stat.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		stats = append(stats, &stat)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return stats, nil
}

func (repo *statisticRepository) ListByDate(date string) ([]*model.TaskStatistic, error) {
	query := `
		SELECT id, milliseconds, task_id, created_at, updated_at
		FROM stats
		WHERE created_at::date = $1;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := repo.db.QueryContext(ctx, query, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats []*model.TaskStatistic

	for rows.Next() {
		var stat model.TaskStatistic

		err := rows.Scan(
			&stat.ID,
			&stat.Milliseconds,
			&stat.TaskID,
			&stat.CreatedAt,
			&stat.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		stats = append(stats, &stat)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return stats, nil
}

func (repo *statisticRepository) Get(id int) (*model.TaskStatistic, error) {
	query := `
		SELECT id, milliseconds, task_id, created_at, updated_at
		FROM stats
		WHERE id = $1;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var stat model.TaskStatistic

	err := repo.db.QueryRowContext(ctx, query, id).Scan(
		&stat.ID,
		&stat.Milliseconds,
		&stat.TaskID,
		&stat.CreatedAt,
		&stat.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &stat, nil
}

func (repo *statisticRepository) Update(stat *model.TaskStatistic) (*model.TaskStatistic, error) {
	query := `
		UPDATE stats 
		SET milliseconds = $1
		WHERE id = $3
		RETURNING id, milliseconds, task_id, created_at, updated_at;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []interface{}{stat.Milliseconds, stat.ID}

	err := repo.db.QueryRowContext(ctx, query, args...).Scan(
		&stat.ID,
		&stat.Milliseconds,
		&stat.TaskID,
		&stat.CreatedAt,
		&stat.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return stat, nil
}

func (repo *statisticRepository) TaskTotalTime(taskId int) (*int, error) {
	query := `
		SELECT coalesce(sum(milliseconds), 0)
		FROM stats
		WHERE task_id = $1;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var sum int

	err := repo.db.QueryRowContext(ctx, query, taskId).Scan(&sum)
	if err != nil {
		return nil, err
	}
	return &sum, nil
}

func (repo *statisticRepository) TaskTodayTime(taskId int) (*int, error) {
	query := `
		SELECT coalesce(sum(milliseconds), 0)
		FROM stats
		WHERE task_id = $1  AND created_at::date = $2;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var sum int

	currentTime := time.Now().Format("2006-01-02")

	err := repo.db.QueryRowContext(ctx, query, taskId, currentTime).Scan(&sum)
	if err != nil {
		return nil, err
	}

	return &sum, nil
}

func (repo *statisticRepository) TaskTotalTimeForDay(taskId int, day string) (int, error) {
	query := `
		SELECT coalesce(sum(milliseconds), 0)
		FROM stats
		WHERE task_id = $1  AND created_at::date = $2;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var sum int

	date, err := time.Parse("2006-01-02", day)
	if err != nil {
		return 0, err
	}

	err = repo.db.QueryRowContext(ctx, query, taskId, date).Scan(&sum)
	if err != nil {
		return 0, err
	}

	return sum, nil
}

func (repo *statisticRepository) TaskTotalTimeFor(taskId int, days int, hours int) (int, error) {
	query := `
		SELECT coalesce(sum(milliseconds), 0)
		FROM stats
		WHERE task_id = $1 AND created_at >= now() - ($2 * interval '1 day' + $3 * interval '1 hour');
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var sum int

	err := repo.db.QueryRowContext(ctx, query, taskId, days, hours).Scan(&sum)
	if err != nil {
		return 0, err
	}

	return sum, nil
}

func (repo *statisticRepository) LastRecord(taskId int) (*model.TaskStatistic, error) {
	query := `
		SELECT id, milliseconds, task_id, created_at, updated_at
		FROM stats
		WHERE task_id = $1 AND created_at::date = $2
		ORDER BY created_at DESC
		LIMIT 1;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var stat model.TaskStatistic

	currentTime := time.Now().Format("2006-01-02")

	err := repo.db.QueryRowContext(ctx, query, taskId, currentTime).Scan(
		&stat.ID,
		&stat.Milliseconds,
		&stat.TaskID,
		&stat.CreatedAt,
		&stat.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &stat, nil
}
