package handler

import (
	"net/http"
	"task-master-api/internal/assignment/application/dtos"
	"task-master-api/internal/assignment/domain/enums"
	"task-master-api/internal/assignment/domain/interfaces"

	commonInterfaces "task-master-api/internal/common/domain/interfaces"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
)

type assignHandler struct {
	service interfaces.AssignService
}

type ResultHandler struct {
	fx.Out

	Handler commonInterfaces.Handler `group:"public_handlers"`
}

func NewAssignHandler(service interfaces.AssignService) ResultHandler {
	return ResultHandler{
		Handler: &assignHandler{
			service: service,
		},
	}
}

func (a *assignHandler) GetAllAssignments(c *gin.Context) {
	assignments, err := a.service.GetAllAssignments()
	if err != nil {
		log.Error().Err(err).Msg("Error getting all assignments")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, assignments)
}

func (a *assignHandler) GetAssignmentByID(c *gin.Context) {
	assignmentID := c.Param("id")
	assignment, err := a.service.GetAssignmentById(assignmentID)
	if err != nil {
		log.Error().Err(err).Msg("Error getting assignment by ID")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, assignment)
}

func (a *assignHandler) CreateAssignment(c *gin.Context) {
	var assignment dtos.AssignmentDto

	if err := c.ShouldBindJSON(&assignment); err != nil {
		a.handleError(c, err, http.StatusBadRequest, "ErrorBindingJSON")
		return
	}

	createdAssignment, err := a.service.CreateAssignment(&assignment)
	if err != nil {
		a.handleAssignmentError(c, err)
		return
	}

	c.JSON(http.StatusCreated, createdAssignment)
}

func (a *assignHandler) DeleteAssignment(c *gin.Context) {
	id := c.Param("id")

	err := a.service.DeleteAssignment(id)
	if err != nil {
		log.Error().Err(err).Msg("Error deleting assignment")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Assignment deleted successfully"})
}

func (a *assignHandler) UpdateAssignment(c *gin.Context) {
	assignmentID := c.Param("id")

	var assignment dtos.UpdateAssignmentDto

	if err := c.ShouldBindJSON(&assignment); err != nil {
		a.handleError(c, err, http.StatusBadRequest, "ErrorBindingJSON")
		return
	}

	updatedAssignment, err := a.service.UpdateAssignment(assignmentID, &assignment)
	if err != nil {
		log.Error().Err(err).Msg("Error updating assignment")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedAssignment)
}

func (a *assignHandler) GetAllAssignmentsByEmployee(c *gin.Context) {
	employeeID := c.Param("employeeId")

	assignments, tasks, err := a.service.GetAllAssignmentsByEmployee(employeeID)
	if err != nil {
		log.Error().Err(err).Msg("Error getting all assignments by employee")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result := dtos.AssignedTasksDto{
		Assignment: assignments,
		Tasks:      tasks,
	}

	c.JSON(http.StatusOK, result)
}

func (a *assignHandler) handleError(c *gin.Context, err error, code int, message string) {
	log.Error().Err(err).Msg(message)
	c.JSON(code, gin.H{"error": err.Error()})
}

func (a *assignHandler) handleAssignmentError(c *gin.Context, err error) {
	switch err.Error() {
	case string(enums.TaskAlreadyCanceled):
		a.handleError(c, err, http.StatusConflict, "TaskAlreadyCanceled")
	case string(enums.TaskAlreadyCompleted):
		a.handleError(c, err, http.StatusConflict, "TaskAlreadyCompleted")
	case string(enums.TaskAlreadyInProgress):
		a.handleError(c, err, http.StatusConflict, "TaskAlreadyInProgress")
	case string(enums.NoEmployeeAvailable):
		a.handleError(c, err, http.StatusNotFound, "NoEmployeeAvailable")
	default:
		a.handleError(c, err, http.StatusInternalServerError, "ErrorCreatingAssignment")
	}
}

func (a *assignHandler) RegisterRoutes(g *gin.RouterGroup) {
	routes := g.Group("/assignment")
	routes.GET("all", a.GetAllAssignments)
	routes.GET(":id", a.GetAssignmentByID)
	routes.POST("", a.CreateAssignment)
	routes.DELETE(":id", a.DeleteAssignment)
	routes.PUT(":id", a.UpdateAssignment)
	routes.GET("employee/:employeeId", a.GetAllAssignmentsByEmployee)
}
