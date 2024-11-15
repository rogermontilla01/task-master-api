package main

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
)

func setLifeCycle(params Params) {
	params.Lc.Append(fx.Hook{
		OnStart: func(context.Context) error {

			registerPublicRoutes(&params)

			go func() {
				err := params.Gin.Run(fmt.Sprintf(":%s", params.Config.Port))
				if err != nil {
					log.Error().Err(err).Msg("Failed to start Gin server")
				}
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			return nil
		},
	})
}

func registerPublicRoutes(p *Params) {
	for _, h := range p.PublicHandlers {
		router := p.Gin.RouterGroup
		h.RegisterRoutes(&router)
	}
}
