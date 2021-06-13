//go:generate go run github.com/99designs/gqlgen
package graph

import (
	"github.com/marsel1323/timetrackerapi/service"
)

type Resolver struct {
	taskService      service.ITaskService
	statisticService service.IStatisticService
	categoryService  service.ICategoryService
}

func NewResolver(
	taskService service.ITaskService,
	statisticService service.IStatisticService,
	categoryService service.ICategoryService,
) *Resolver {
	return &Resolver{
		taskService:      taskService,
		statisticService: statisticService,
		categoryService:  categoryService,
	}
}
