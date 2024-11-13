package config

import "go.uber.org/fx"

var Module = fx.Module("appCofig", fx.Provide(
	NewApiConfig,
))
