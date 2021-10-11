package repository

import (
	"context"
	"database/sql"
	"github.com/marsel1323/timetrackerapi/graph/model"
	"time"
)

type GoalRepository interface {
	Create(goal model.NewGoal) (*model.Goal, error)
	List() ([]*model.Goal, error)
}

type goalRepository struct {
	db *sql.DB
}

func NewGoalRepository(db *sql.DB) *goalRepository {
	return &goalRepository{
		db: db,
	}
}

func (repo *goalRepository) Create(input model.NewGoal) (*model.Goal, error) {
	query := `
		INSERT INTO goals (name, time, category_id)
		VALUES ($1, $2, $3)
		RETURNING id, name, time, category_id, created_at, updated_at;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []interface{}{input.Name, input.Time, input.CategoryID}

	var goal model.Goal

	err := repo.db.QueryRowContext(ctx, query, args...).
		Scan(
			&goal.ID,
			&goal.Name,
			&goal.Time,
			&goal.CategoryID,
			&goal.CreatedAt,
			&goal.UpdatedAt,
		)
	if err != nil {
		return nil, err
	}

	return &goal, nil
}

func (repo *goalRepository) List() ([]*model.Goal, error) {
	query := `
		SELECT id, name, time, category_id, created_at, updated_at
		FROM goals
		ORDER BY id;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var goals []*model.Goal

	for rows.Next() {
		var goal model.Goal

		err := rows.Scan(
			&goal.ID,
			&goal.Name,
			&goal.Time,
			&goal.CategoryID,
			&goal.CreatedAt,
			&goal.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		goals = append(goals, &goal)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return goals, nil
}
