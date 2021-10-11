//go:generate go run github.com/99designs/gqlgen
package graph

import (
	"github.com/marsel1323/timetrackerapi/service"
)

type Resolver struct {
	taskService      service.TaskService
	statisticService service.IStatisticService
	categoryService  service.ICategoryService
	goalService      service.GoalService
	goalStatisticService service.GoalStatisticService
}

func NewResolver(
	taskService service.TaskService,
	statisticService service.IStatisticService,
	categoryService service.ICategoryService,
	goalService service.GoalService,
	goalStatisticService service.GoalStatisticService,
) *Resolver {
	return &Resolver{
		taskService:      taskService,
		statisticService: statisticService,
		categoryService:  categoryService,
		goalService:      goalService,
		goalStatisticService: goalStatisticService,
	}
}
