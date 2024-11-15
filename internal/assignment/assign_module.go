package assignment

import (
	"task-master-api/internal/assignment/application/service"
	"task-master-api/internal/assignment/domain/interfaces"
	"task-master-api/internal/assignment/infrastructure/handler"
	"task-master-api/internal/assignment/infrastructure/repository"

	"go.uber.org/fx"
)

var Module = fx.Module("assignment", fx.Provide(
	handler.NewAssignHandler,
	fx.Annotate(
		service.NewAssignmentService,
		fx.As(new(interfaces.AssignService)),
	),
	fx.Annotate(
		repository.NewAssignmentRepository,
		fx.As(new(interfaces.AssignmentRepository)),
	),
))
