package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/marsel1323/timetrackerapi/graph/generated"
	"github.com/marsel1323/timetrackerapi/graph/model"
)

func (r *categoryResolver) TotalTime(ctx context.Context, obj *model.Category) (int, error) {
	result, err := r.categoryService.TotalTime(obj.ID)
	if err != nil {
		//log.Error(err)
		return 0, nil
	}
	return result, nil
}

func (r *categoryResolver) TodayTime(ctx context.Context, obj *model.Category) (int, error) {
	result, err := r.categoryService.TodayTime(obj.ID)
	if err != nil {
		//log.Error(err)
		return 0, nil
	}
	return result, nil
}

func (r *goalResolver) TodayMs(ctx context.Context, obj *model.Goal) (int, error) {
	return r.goalStatisticService.TodayTime(obj.ID)
}

func (r *mutationResolver) CreateCategory(ctx context.Context, input model.NewCategory) (*model.Category, error) {
	return r.categoryService.CreateCategory(input)
}

func (r *mutationResolver) CreateTask(ctx context.Context, input model.NewTask) (*model.Task, error) {
	task, err := r.taskService.CreateTask(&input)
	if err != nil {
		log.Println(err)
	}
	return task, err
}

func (r *mutationResolver) CreateTaskStatistic(ctx context.Context, input model.NewTaskStatistic) (*model.TaskStatistic, error) {
	return r.statisticService.CreateStat(input)
}

func (r *mutationResolver) CreateGoal(ctx context.Context, input model.NewGoal) (*model.Goal, error) {
	return r.goalService.CreateGoal(input)
}

func (r *queryResolver) CategoriesList(ctx context.Context) ([]*model.Category, error) {
	return r.categoryService.List()
}

func (r *queryResolver) Task(ctx context.Context, id int) (*model.Task, error) {
	return r.taskService.GetTask(id)
}

func (r *queryResolver) TaskList(ctx context.Context) ([]*model.Task, error) {
	taskList, err := r.taskService.TaskList()
	if err != nil {
		log.Println(err)
	}
	return taskList, err
}

func (r *queryResolver) TaskListByCategory(ctx context.Context, categoryID int) ([]*model.Task, error) {
	taskList, err := r.taskService.TaskListByCategory(categoryID)
	if err != nil {
		log.Println(err)
	}
	return taskList, err
}

func (r *queryResolver) Goal(ctx context.Context, id int) (*model.Goal, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GoalList(ctx context.Context) ([]*model.Goal, error) {
	return r.goalService.GoalList()
}

func (r *queryResolver) StatListByDate(ctx context.Context, date string) ([]*model.TaskStatistic, error) {
	return r.statisticService.StatListByDate(date)
}

func (r *taskResolver) Category(ctx context.Context, obj *model.Task) (*model.Category, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *taskResolver) TotalMs(ctx context.Context, obj *model.Task) (*int, error) {
	return r.statisticService.CalcTotalTime(obj.ID)
}

func (r *taskResolver) TotalToday(ctx context.Context, obj *model.Task) (*int, error) {
	return r.statisticService.CalcTotalTodayTime(obj.ID)
}

func (r *taskResolver) TotalTimeFor(ctx context.Context, obj *model.Task, day string) (int, error) {
	return r.statisticService.TotalTimeFor(obj.ID, day)
}

func (r *taskResolver) TotalTimeForLast(ctx context.Context, obj *model.Task, days int, hours int) (int, error) {
	return r.statisticService.CalcTotalTimeFor(obj.ID, days, hours)
}

func (r *taskResolver) LastStat(ctx context.Context, obj *model.Task) (*model.TaskStatistic, error) {
	data, err := r.statisticService.LastStatRecord(obj.ID)
	if err != nil {
		log.Println(err)
		return nil, nil
	}
	return data, nil
}

// Category returns generated.CategoryResolver implementation.
func (r *Resolver) Category() generated.CategoryResolver { return &categoryResolver{r} }

// Goal returns generated.GoalResolver implementation.
func (r *Resolver) Goal() generated.GoalResolver { return &goalResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Task returns generated.TaskResolver implementation.
func (r *Resolver) Task() generated.TaskResolver { return &taskResolver{r} }

type categoryResolver struct{ *Resolver }
type goalResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type taskResolver struct{ *Resolver }
