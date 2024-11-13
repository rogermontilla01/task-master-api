package interfaces

import "task-master-api/internal/employee/application/dtos"

type EmployeeRepository interface {
	CreateEmployee(employee *dtos.EmployeeDto) (*dtos.EmployeeDto, error)
}
