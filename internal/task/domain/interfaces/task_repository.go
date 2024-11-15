package interfaces

import "task-master-api/internal/task/application/dtos"

type TaskRepository interface {
	GetAllTasks() (*[]dtos.TaskDto, error)
	GetTaskById(id string) (*dtos.TaskDto, error)
	CreateTask(task *dtos.TaskDto) (*dtos.TaskDto, error)
	UpdateTask(id string, task *dtos.UpdateTaskDto) (*dtos.UpdateTaskDto, error)
	DeleteTask(id string) error
	GetAllTasksById(id string) (*[]dtos.TaskDto, error)
}
