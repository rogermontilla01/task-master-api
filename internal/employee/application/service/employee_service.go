package service

import (
	"task-master-api/internal/employee/application/dtos"
	"task-master-api/internal/employee/domain/interfaces"

	"github.com/rs/zerolog/log"
)

type EmployeeService struct {
	repository interfaces.EmployeeRepository
}

func NewEmployeeService(
	repository interfaces.EmployeeRepository,
) interfaces.EmployeeService {
	return &EmployeeService{
		repository: repository,
	}
}

func (e *EmployeeService) GetAllEmployees() (*[]dtos.EmployeeDto, error) {
	log.Info().Msg("Start getting employees")

	employees, err := e.repository.GetAllEmployees()
	if err != nil {
		log.Error().Err(err).Msg("error getting employees")
		return nil, err
	}

	return employees, nil
}

func (e *EmployeeService) GetEmployeeById(id string) (employee *dtos.EmployeeDto, err error) {
	log.Info().Str("id", id).Msg("Get employee")

	employeeDto, err := e.repository.GetEmployee(id)
	if err != nil {
		log.Error().Err(err).Msg("error getting employee")
		return employeeDto, err
	}

	return employeeDto, nil
}

func (e *EmployeeService) CreateEmployee(employee *dtos.EmployeeDto) (*dtos.EmployeeDto, error) {
	log.Info().Msg("Start employee creation")

	employeeDto, err := e.repository.CreateEmployee(employee)
	if err != nil {
		log.Error().Err(err).Msg("error creating employee")
		return employeeDto, err
	}

	log.Info().Msg("employee created")

	return employeeDto, nil

}

func (e *EmployeeService) UpdateEmployee(id string, employee *dtos.UpdateEmployeeDto) (*dtos.UpdateEmployeeDto, error) {

	_, err := e.repository.UpdateEmployee(id, employee)
	if err != nil {
		log.Error().Err(err).Msg("error updating employee")
		return nil, err
	}

	return employee, nil
}

func (e *EmployeeService) DeleteEmployee(id string) error {
	log.Info().Str("id", id).Msg("Start employee deletion")

	err := e.repository.DeleteEmployee(id)
	if err != nil {
		log.Error().Err(err).Msg("error deleting employee")
		return err
	}

	log.Info().Msg("employee deleted")

	return nil
}
