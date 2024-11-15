package interfaces

import (
	"task-master-api/internal/assignment/application/dtos"
	taskDtos "task-master-api/internal/task/application/dtos"
)

type AssignService interface {
	GetAllAssignments() (*[]dtos.AssignmentDto, error)
	GetAssignmentById(assignId string) (*dtos.AssignmentDto, error)
	CreateAssignment(*dtos.AssignmentDto) (*dtos.AssignmentDto, error)
	DeleteAssignment(assignId string) error
	UpdateAssignment(assignId string, update *dtos.UpdateAssignmentDto) (*dtos.UpdateAssignmentDto, error)
	GetAllAssignmentsByEmployee(employeeId string) (*[]dtos.AssignmentDto, *[]taskDtos.TaskDto, error)
}
