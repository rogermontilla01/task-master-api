package mongo

import (
	"context"
	"task-master-api/internal/config"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	URL   string `env:"value=MONGO_URL,required"`
	DB    string `env:"value=MONGO_DATABASE,required"`
	Debug bool   `env:"value=MONGO_DEBUG,default=false"`
}

type queryLogger struct{}

func (ql queryLogger) Started(ctx context.Context, evt *event.CommandStartedEvent) {
	log.Debug().Msgf("Command started: %s %d %s", evt.CommandName, evt.RequestID, evt.Command)
}

func (ql queryLogger) Succeeded(ctx context.Context, evt *event.CommandSucceededEvent) {
	log.Debug().Msgf("Command succeeded: %s %d", evt.CommandName, evt.RequestID)
}

func (ql queryLogger) Failed(ctx context.Context, evt *event.CommandFailedEvent) {
	log.Debug().Msgf("Command failed: %s %d %s", evt.CommandName, evt.RequestID, evt.Failure)
}

func NewMongoClient() *mongo.Database {

	mongoConfig := MongoConfig{}

	config.ValidateEnvConfigOrFail(&mongoConfig)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoConfig.URL)

	if mongoConfig.Debug {
		monitor := &event.CommandMonitor{
			Started:   queryLogger{}.Started,
			Succeeded: queryLogger{}.Succeeded,
			Failed:    queryLogger{}.Failed,
		}
		clientOptions.SetMonitor(monitor)
	}

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal().Err(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal().Err(err)
	}

	log.Info().Msg("Connected to mongo successfully")

	return client.Database(mongoConfig.DB)
}
