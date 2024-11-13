package handler

import (
	"net/http"
	commonInterfaces "task-master-api/internal/common/domain/interfaces"
	"task-master-api/internal/employee/application/dtos"
	"task-master-api/internal/employee/domain/interfaces"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
)

type employeeHandler struct {
	service interfaces.EmployeeService
}

type ResultHandler struct {
	fx.Out

	Handler commonInterfaces.Handler `group:"public_handlers"`
}

func NewCronjobHandler(service interfaces.EmployeeService) ResultHandler {
	return ResultHandler{
		Handler: &employeeHandler{
			service: service,
		},
	}
}

func (h *employeeHandler) CreateEmployee(c *gin.Context) {
	var employee dtos.EmployeeDto
	if err := c.ShouldBindBodyWith(&employee, binding.JSON); err != nil {
		log.Error().Caller().Err(err).Send()
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newEmployee, err := h.service.CreateEmployee(&employee)

	if err != nil {
		log.Error().Caller().Err(err).Send()
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newEmployee)
}

func (h *employeeHandler) GetEmployee(c *gin.Context) {
	id := h.service.GetEmployee("test")

	c.JSON(http.StatusOK, id)
}

func (h *employeeHandler) RegisterRoutes(g *gin.RouterGroup) {
	routes := g.Group("/employee")
	routes.GET("/:id", h.GetEmployee)
	routes.POST("", h.CreateEmployee)
}
