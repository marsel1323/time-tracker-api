package repository

import (
	"context"
	"database/sql"
	"time"
)

type GoalStatisticRepository interface {
	TodayTime(goalId int) (int, error)
}

type goalStatisticRepository struct {
	db *sql.DB
}

func NewGoalStatisticRepository(db *sql.DB) *goalStatisticRepository {
	return &goalStatisticRepository{
		db: db,
	}
}

func (repo *goalStatisticRepository) TodayTime(goalId int) (int, error) {
	query := `
		
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var sum = 0
	today := time.Now().Format("2006-01-02")

	err := repo.db.QueryRowContext(ctx, query, goalId, today).Scan(&sum)
	if err != nil {
		return 0, err
	}

	return sum, nil
}
