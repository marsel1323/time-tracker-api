package repository

import (
	"context"
	"database/sql"
	"github.com/marsel1323/timetrackerapi/graph/model"
	"time"
)

type IStatisticRepository interface {
	Create(stat model.Statistic) (*model.Statistic, error)
	List() ([]*model.Statistic, error)
	ListByDate(date string) ([]*model.Statistic, error)
	Get(id int) (*model.Statistic, error)
	Update(stat *model.Statistic) (*model.Statistic, error)

	TaskTotalTime(taskId int) (*int, error)
	TaskTodayTime(taskId int) (*int, error)
	TaskTotalTimeFor(taskId int, days int, hours int) (int, error)
	TaskTotalTimeForDay(taskId int, day string) (int, error)
	LastRecord(taskId int) (*model.Statistic, error)
}

type StatisticRepository struct {
	db *sql.DB
}

func NewStatisticRepository(db *sql.DB) *StatisticRepository {
	return &StatisticRepository{
		db: db,
	}
}

func (repo *StatisticRepository) Create(stat model.Statistic) (*model.Statistic, error) {
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

func (repo *StatisticRepository) List() ([]*model.Statistic, error) {
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

	var stats []*model.Statistic

	for rows.Next() {
		var stat model.Statistic

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

func (repo *StatisticRepository) ListByDate(date string) ([]*model.Statistic, error) {
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

	var stats []*model.Statistic

	for rows.Next() {
		var stat model.Statistic

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

func (repo *StatisticRepository) Get(id int) (*model.Statistic, error) {
	query := `
		SELECT id, milliseconds, task_id, created_at, updated_at
		FROM stats
		WHERE id = $1;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var stat model.Statistic

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

func (repo *StatisticRepository) Update(stat *model.Statistic) (*model.Statistic, error) {
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

func (repo *StatisticRepository) TaskTotalTime(taskId int) (*int, error) {
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

func (repo *StatisticRepository) TaskTodayTime(taskId int) (*int, error) {
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

func (repo *StatisticRepository) TaskTotalTimeForDay(taskId int, day string) (int, error) {
	query := `
		SELECT coalesce(sum(milliseconds), 0)
		FROM stats
		WHERE task_id = $1  AND created_at::date = $2;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var sum int

	// currentTime := time.Now().Format("2006-01-02")
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

func (repo *StatisticRepository) TaskTotalTimeFor(taskId int, days int, hours int) (int, error) {
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

func (repo *StatisticRepository) LastRecord(taskId int) (*model.Statistic, error) {
	query := `
		SELECT id, milliseconds, task_id, created_at, updated_at
		FROM stats
		WHERE task_id = $1 AND created_at::date = $2
		ORDER BY created_at DESC
		LIMIT 1;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var stat model.Statistic

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
