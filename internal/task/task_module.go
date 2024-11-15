package task

import (
	"task-master-api/internal/task/application/service"
	"task-master-api/internal/task/domain/interfaces"
	"task-master-api/internal/task/infrastructure/handler"
	"task-master-api/internal/task/infrastructure/repository"

	"go.uber.org/fx"
)

var Module = fx.Module("task", fx.Provide(
	handler.NewTaskHandler,
	fx.Annotate(
		service.NewTaskService,
		fx.As(new(interfaces.TaskService)),
	),
	fx.Annotate(
		repository.NewTaskRepository,
		fx.As(new(interfaces.TaskRepository)),
	),
))
