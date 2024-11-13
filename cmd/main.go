package main

import (
	"task-master-api/internal/common"
	"task-master-api/internal/common/domain/interfaces"
	"task-master-api/internal/config"
	"task-master-api/internal/employee"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Gin            *gin.Engine
	Config         config.ApiConfig
	Lc             fx.Lifecycle
	PublicHandlers []interfaces.Handler `group:"public_handlers"`
}

func main() {

	app := fx.New(
		config.Module,
		employee.Module,
		common.Module,

		fx.Provide(
			gin.Default,
		),

		fx.Invoke(
			setLifeCycle,
		),
	)

	app.Run()
}
