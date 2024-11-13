package employee

import (
	"task-master-api/internal/employee/application/service"
	"task-master-api/internal/employee/domain/interfaces"
	"task-master-api/internal/employee/infrastructure/handler"
	"task-master-api/internal/employee/infrastructure/repository"

	"go.uber.org/fx"
)

var Module = fx.Module("employee", fx.Provide(
	handler.NewCronjobHandler,
	fx.Annotate(
		service.NewEmployeeService,
		fx.As(new(interfaces.EmployeeService)),
	),
	fx.Annotate(
		repository.NewEmployeeRepository,
		fx.As(new(interfaces.EmployeeRepository)),
	),
))
