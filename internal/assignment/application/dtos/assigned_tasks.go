package dtos

import (
	taskDtos "task-master-api/internal/task/application/dtos"
)

type AssignedTasksDto struct {
	Tasks      *[]taskDtos.TaskDto `json:"tasks"`
	Assignment *[]AssignmentDto    `json:"assignment"`
}
