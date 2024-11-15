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

func (e *employeeHandler) CreateEmployee(c *gin.Context) {
	var employee dtos.EmployeeDto
	if err := c.ShouldBindBodyWith(&employee, binding.JSON); err != nil {
		log.Error().Caller().Err(err).Send()
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newEmployee, err := e.service.CreateEmployee(&employee)
	if err != nil {
		log.Error().Caller().Err(err).Send()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newEmployee)
}

func (e *employeeHandler) GetEmployee(c *gin.Context) {
	id := c.Param("id")

	employee, err := e.service.GetEmployee(id)
	if err != nil {
		log.Error().Caller().Err(err).Send()
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employee)
}

func (e *employeeHandler) UpdateEmployee(c *gin.Context) {
	id := c.Param("id")

	var employee dtos.UpdateEmployeeDto
	if err := c.ShouldBindBodyWith(&employee, binding.JSON); err != nil {
		log.Error().Caller().Err(err).Send()
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Employee data: %+v\n", employee)

	updatedEmployee, err := e.service.UpdateEmployee(id, &employee)
	if err != nil {
		log.Error().Caller().Err(err).Send()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedEmployee)
}

func (e *employeeHandler) DeleteEmployee(c *gin.Context) {
	id := c.Param("id")

	err := e.service.DeleteEmployee(id)
	if err != nil {
		log.Error().Caller().Err(err).Send()
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

func (e *employeeHandler) RegisterRoutes(g *gin.RouterGroup) {
	routes := g.Group("/employee")
	routes.GET("/:id", e.GetEmployee)
	routes.POST("", e.CreateEmployee)
	routes.PUT("/:id", e.UpdateEmployee)
	routes.DELETE("/:id", e.DeleteEmployee)
}
