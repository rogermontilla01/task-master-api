package common

import (
	mongo "task-master-api/internal/common/infrastructure"

	"go.uber.org/fx"
)

var Module = fx.Module("mongodb", fx.Provide(
	mongo.NewMongoClient,
))
