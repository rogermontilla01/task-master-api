package service

import (
	"task-master-api/internal/task/application/dtos"
	"task-master-api/internal/task/domain/interfaces"

	"github.com/rs/zerolog/log"
)

type TaskService struct {
	repository interfaces.TaskRepository
}

func NewTaskService(repository interfaces.TaskRepository) interfaces.TaskService {
	return &TaskService{repository: repository}
}

func (t *TaskService) CreateTask(task *dtos.TaskDto) (*dtos.TaskDto, error) {
	log.Info().Msg("Creating task")

	taskDto, err := t.repository.CreateTask(task)
	if err != nil {
		log.Error().Err(err).Msg("Error creating task")
		return nil, err
	}

	log.Info().Msg("Task created successfully")

	return taskDto, nil
}

func (t *TaskService) DeleteTask(id string) error {
	log.Info().Msg("Deleting task")

	err := t.repository.DeleteTask(id)
	if err != nil {
		log.Error().Err(err).Msg("Error deleting task")
		return err
	}

	log.Info().Msg("Task deleted successfully")

	return nil
}

func (t *TaskService) GetAllTasks() (*[]dtos.TaskDto, error) {
	log.Info().Msg("Getting all tasks")

	tasks, err := t.repository.GetAllTasks()
	if err != nil {
		log.Error().Err(err).Msg("Error getting all tasks")
		return nil, err
	}

	log.Info().Msg("Tasks retrieved successfully")

	return tasks, nil
}

func (t *TaskService) GetTaskById(id string) (*dtos.TaskDto, error) {
	log.Info().Msg("Getting task by id")

	task, err := t.repository.GetTaskById(id)
	if err != nil {
		log.Error().Err(err).Msg("Error getting task by id")
		return nil, err
	}

	log.Info().Msg("Task retrieved successfully")

	return task, nil
}

func (t *TaskService) UpdateTask(id string, task *dtos.UpdateTaskDto) (*dtos.UpdateTaskDto, error) {
	log.Info().Msg("Updating task")

	updatedTask, err := t.repository.UpdateTask(id, task)
	if err != nil {
		log.Error().Err(err).Msg("Error updating task")
		return nil, err
	}

	log.Info().Msg("Task updated successfully")

	return updatedTask, nil
}
