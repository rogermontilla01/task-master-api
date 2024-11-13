package service

import (
	"fmt"
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

func (e *EmployeeService) GetEmployee(id string) string {
	fmt.Println("call form employee service")

	return id
}

func (e *EmployeeService) CreateEmployee(employee *dtos.EmployeeDto) (*dtos.EmployeeDto, error) {
	log.Info().Msg("Start employee creation")

	dto, err := e.repository.CreateEmployee(employee)
	if err != nil {
		log.Error().Err(err).Msg("error creating employee")
		return dto, err
	}

	log.Info().Msg("employee created")

	return dto, nil

}
