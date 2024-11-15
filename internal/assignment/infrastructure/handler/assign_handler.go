package handler

import (
	"net/http"
	"task-master-api/internal/assignment/application/dtos"
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
		log.Error().Err(err).Msg("Error binding JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdAssignment, err := a.service.CreateAssignment(&assignment)
	if err != nil {
		log.Error().Err(err).Msg("Error creating assignment")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
		log.Error().Err(err).Msg("Error binding JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

func (a *assignHandler) RegisterRoutes(g *gin.RouterGroup) {
	routes := g.Group("/assignment")
	routes.GET("all", a.GetAllAssignments)
	routes.GET(":id", a.GetAssignmentByID)
	routes.POST("", a.CreateAssignment)
	routes.DELETE(":id", a.DeleteAssignment)
	routes.PUT(":id", a.UpdateAssignment)
}
