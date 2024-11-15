package interfaces

import "github.com/gin-gonic/gin"

type Handler interface {
	RegisterRoutes(router *gin.RouterGroup)
}
