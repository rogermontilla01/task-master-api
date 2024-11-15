package interfaces

import "task-master-api/internal/employee/application/dtos"

type EmployeeRepository interface {
	GetEmployee(id string) (*dtos.EmployeeDto, error)
	CreateEmployee(employee *dtos.EmployeeDto) (*dtos.EmployeeDto, error)
	UpdateEmployee(id string, employee *dtos.UpdateEmployeeDto) (*dtos.UpdateEmployeeDto, error)
	DeleteEmployee(id string) error
}
