package service

import (
	"task-master-api/internal/assignment/application/dtos"
	"task-master-api/internal/assignment/domain/interfaces"
	employeesDtos "task-master-api/internal/employee/application/dtos"
	employeeInterfaces "task-master-api/internal/employee/domain/interfaces"
	taskInterfaces "task-master-api/internal/task/domain/interfaces"
	"time"

	"github.com/rs/zerolog/log"
)

type AssignmentService struct {
	repository      interfaces.AssignmentRepository
	taskService     taskInterfaces.TaskService
	employeeService employeeInterfaces.EmployeeService
}

func NewAssignmentService(
	repository interfaces.AssignmentRepository,
	taskService taskInterfaces.TaskService,
	employeeService employeeInterfaces.EmployeeService,
) interfaces.AssignService {
	return &AssignmentService{
		repository,
		taskService,
		employeeService,
	}
}

func (a *AssignmentService) CreateAssignment(assignment *dtos.AssignmentDto) (*dtos.AssignmentDto, error) {
	log.Info().Msg("creating assignment")

	task, err := a.taskService.GetTaskById(assignment.TaskID)
	if err != nil {
		log.Error().Err(err).Msg("error getting task")
		return nil, err
	}

	employees, err := a.employeeService.GetAllEmployees()
	if err != nil {
		log.Error().Err(err).Msg("error getting employee")
		return nil, err
	}

	employeesBySkill := a.filterEmployeesBySkills(*employees, task.Skills)

	employeeByAvailableTime := a.findEmployeeWithAvailableTime(employeesBySkill, task.Duration)

	if employeeByAvailableTime == nil {
		log.Error().Msg("no employee available")
		return nil, err
	}

	log.Info().Interface("employee", employeeByAvailableTime).Msg("Empleado registrado")

	//TODO: continue here

	return nil, nil
}

func (a *AssignmentService) findEmployeeWithAvailableTime(employees []employeesDtos.EmployeeDto, taskDuration string) *employeesDtos.EmployeeDto {
	for _, employee := range employees {
		availableTime, err := time.ParseDuration(employee.AvailableHours)
		if err != nil {
			continue
		}
		taskDurationParsed, err := time.ParseDuration(taskDuration)
		if err != nil {
			continue
		}
		if availableTime >= taskDurationParsed {
			return &employee
		}
	}
	return nil
}

func (a *AssignmentService) filterEmployeesBySkills(employees []employeesDtos.EmployeeDto, requiredSkills []string) []employeesDtos.EmployeeDto {
	var filteredEmployees []employeesDtos.EmployeeDto

	for _, employee := range employees {
		if a.hasRequiredSkills(employee.Skills, requiredSkills) {
			filteredEmployees = append(filteredEmployees, employee)
		}
	}

	return filteredEmployees
}

func (a *AssignmentService) hasRequiredSkills(employeeSkills []string, taskSkills []string) bool {
	for _, skill := range taskSkills {
		if !a.contains(employeeSkills, skill) {
			return false
		}
	}
	return true
}

func (a *AssignmentService) contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func (a *AssignmentService) DeleteAssignment(assignId string) error {
	panic("unimplemented")
}

func (a *AssignmentService) GetAllAssignments() (*[]dtos.AssignmentDto, error) {
	panic("unimplemented")
}

func (a *AssignmentService) GetAssignmentById(assignId string) (*dtos.AssignmentDto, error) {
	panic("unimplemented")
}

func (a *AssignmentService) UpdateAssignment(assignId string, update *dtos.UpdateAssignmentDto) (*dtos.UpdateAssignmentDto, error) {
	panic("unimplemented")
}
