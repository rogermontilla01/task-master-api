package interfaces

import (
	"task-master-api/internal/employee/application/dtos"
)

type EmployeeService interface {
	GetEmployee(id string) string
	CreateEmployee(employee *dtos.EmployeeDto) (*dtos.EmployeeDto, error)
}
