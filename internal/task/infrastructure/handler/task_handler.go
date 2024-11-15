package handler

import (
	"net/http"
	commonInterfaces "task-master-api/internal/common/domain/interfaces"
	"task-master-api/internal/task/application/dtos"
	"task-master-api/internal/task/domain/interfaces"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
)

type taskHandler struct {
	service interfaces.TaskService
}

type ResultHandler struct {
	fx.Out

	Handler commonInterfaces.Handler `group:"public_handlers"`
}

func NewTaskHandler(service interfaces.TaskService) ResultHandler {
	return ResultHandler{
		Handler: &taskHandler{
			service: service,
		},
	}
}

func (e *taskHandler) GetAllTasks(c *gin.Context) {
	tasks, err := e.service.GetAllTasks()
	if err != nil {
		log.Error().Err(err).Msg("Error getting all tasks")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (e *taskHandler) GetTaskByID(c *gin.Context) {
	id := c.Param("id")

	task, err := e.service.GetTaskById(id)
	if err != nil {
		log.Error().Err(err).Msg("Error getting task by ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (e *taskHandler) CreateTask(c *gin.Context) {
	var task dtos.TaskDto

	if err := c.ShouldBindJSON(&task); err != nil {
		log.Error().Err(err).Msg("Error binding JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTask, err := e.service.CreateTask(&task)
	if err != nil {
		log.Error().Err(err).Msg("Error creating task")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdTask)
}

func (e *taskHandler) UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var task dtos.UpdateTaskDto

	if err := c.ShouldBindJSON(&task); err != nil {
		log.Error().Err(err).Msg("Error binding JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTask, err := e.service.UpdateTask(id, &task)
	if err != nil {
		log.Error().Err(err).Msg("Error updating task")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedTask)
}

func (e *taskHandler) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	err := e.service.DeleteTask(id)
	if err != nil {
		log.Error().Err(err).Msg("Error deleting task")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

func (e *taskHandler) RegisterRoutes(g *gin.RouterGroup) {
	routes := g.Group("/task")
	routes.GET("all", e.GetAllTasks)
	routes.GET(":id", e.GetTaskByID)
	routes.POST("", e.CreateTask)
	routes.PUT(":id", e.UpdateTask)
	routes.DELETE(":id", e.DeleteTask)
}
