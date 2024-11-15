package service

import (
	"errors"
	"task-master-api/internal/assignment/application/dtos"
	"task-master-api/internal/assignment/domain/enums"
	"task-master-api/internal/assignment/domain/interfaces"
	employeesDtos "task-master-api/internal/employee/application/dtos"
	employeeInterfaces "task-master-api/internal/employee/domain/interfaces"
	taskDtos "task-master-api/internal/task/application/dtos"
	taskEnums "task-master-api/internal/task/domain/enums"
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

	task, err := a.getValidTask(assignment.TaskID)
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

	employeeByAvailable := a.findEmployeeWithAvailableTime(employeesBySkill, *task)

	if employeeByAvailable == nil {
		log.Error().Msg("no employee available")
		return nil, errors.New(string(enums.NoEmployeeAvailable))
	}

	err = a.ReduceAvailableTime(employeeByAvailable, *task)
	if err != nil {
		log.Error().Err(err).Msg("error reducing available time")
		return nil, err
	}

	err = a.setTaskInProgress(task.ID)
	if err != nil {
		log.Error().Err(err).Msg("error setting task in progress")
		return nil, err
	}

	err = a.updateTaskStatus(task.ID)
	if err != nil {
		log.Error().Err(err).Msg("error updating task")
		return nil, err
	}

	assignment.EmployeeID = employeeByAvailable.ID
	assignment.TaskID = task.ID
	assignment.Duration = task.Duration

	newAssignment, err := a.repository.CreateAssignment(assignment)
	if err != nil {
		log.Error().Err(err).Msg("error creating assignment")
		return nil, err
	}

	return newAssignment, nil
}

func (a *AssignmentService) updateTaskStatus(id string) error {
	status := string(taskEnums.InProgress)
	update := taskDtos.UpdateTaskDto{
		Status: &status,
	}
	_, err := a.taskService.UpdateTask(id, &update)
	if err != nil {
		log.Error().Err(err).Msg("error updating task")
		return err
	}

	return nil
}

func (a *AssignmentService) getValidTask(id string) (*taskDtos.TaskDto, error) {
	task, err := a.taskService.GetTaskById(id)
	if err != nil {
		log.Error().Err(err).Msg("error getting task")
		return task, err
	}

	if task.Status == string(taskEnums.InProgress) {
		log.Error().Msg(string(enums.TaskAlreadyInProgress))
		return nil, errors.New(string(enums.TaskAlreadyInProgress))
	}

	if task.Status == string(taskEnums.Completed) {
		log.Error().Msg(string(enums.TaskAlreadyCompleted))
		return nil, errors.New(string(enums.TaskAlreadyCompleted))
	}

	if task.Status == string(taskEnums.Canceled) {
		log.Error().Msg(string(enums.TaskAlreadyCanceled))
		return nil, errors.New(string(enums.TaskAlreadyCanceled))
	}

	return task, nil
}

func (a *AssignmentService) setTaskInProgress(taskId string) error {
	status := "inProgress"
	update := taskDtos.UpdateTaskDto{
		Status: &status,
	}

	_, err := a.taskService.UpdateTask(taskId, &update)

	if err != nil {
		log.Error().Err(err).Msg("error updating task")
		return err
	}

	return nil
}

func (a *AssignmentService) ReduceAvailableTime(employee *employeesDtos.EmployeeDto, task taskDtos.TaskDto) error {
	availableHours, _ := employee.GetAvailableHours()
	duration, _ := task.GetDuration()

	newAvailableHours := availableHours - duration

	parcedDuration := time.Duration(newAvailableHours).String()

	employee.AvailableHours = parcedDuration

	update := employeesDtos.UpdateEmployeeDto{
		AvailableHours: &parcedDuration,
	}

	_, err := a.employeeService.UpdateEmployee(employee.ID, &update)
	if err != nil {
		log.Error().Err(err).Msg("error updating employee")
		return err
	}

	return nil
}

func (a *AssignmentService) findEmployeeWithAvailableTime(employees []employeesDtos.EmployeeDto, task taskDtos.TaskDto) *employeesDtos.EmployeeDto {
	for _, employee := range employees {
		availableTime, err := employee.GetAvailableHours()
		if err != nil {
			return nil
		}

		taskDurationParsed, err := task.GetDuration()
		if err != nil {
			return nil
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
	err := a.repository.DeleteAssignment(assignId)
	if err != nil {
		log.Error().Err(err).Msg("error deleting assignment")
		return err
	}

	return nil
}

func (a *AssignmentService) GetAllAssignments() (*[]dtos.AssignmentDto, error) {
	assignments, err := a.repository.GetAllAssignments()
	if err != nil {
		log.Error().Err(err).Msg("error getting all assignments")
		return nil, err
	}

	return assignments, nil
}

func (a *AssignmentService) GetAssignmentById(assignId string) (*dtos.AssignmentDto, error) {
	assignment, err := a.repository.GetAssignmentById(assignId)
	if err != nil {
		log.Error().Err(err).Msg("error getting assignment")
		return nil, err
	}

	return assignment, nil
}

func (a *AssignmentService) UpdateAssignment(assignId string, update *dtos.UpdateAssignmentDto) (*dtos.UpdateAssignmentDto, error) {
	updatedAssignment, err := a.repository.UpdateAssignment(assignId, update)
	if err != nil {
		log.Error().Err(err).Msg("error updating assignment")
		return nil, err
	}

	return updatedAssignment, nil
}

func (a *AssignmentService) GetAllAssignmentsByEmployee(employeeId string) (*[]dtos.AssignmentDto, *[]taskDtos.TaskDto, error) {
	tasks := []taskDtos.TaskDto{}

	assignments, err := a.repository.GetAllAssignmentsByEmployee(employeeId)
	if err != nil {
		log.Error().Err(err).Msg("error getting all assignments")
		return nil, nil, err
	}

	for _, assignment := range *assignments {
		task, err := a.taskService.GetTaskById(assignment.TaskID)
		if err != nil {
			log.Error().Err(err).Msg("error getting task")
			return nil, nil, err
		}

		tasks = append(tasks, *task)
	}

	return assignments, &tasks, nil
}
