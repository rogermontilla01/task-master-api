package repository

import (
	"context"
	"task-master-api/internal/employee/application/dtos"
	"task-master-api/internal/employee/domain/interfaces"
	"task-master-api/internal/employee/infrastructure/entities"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeRepository struct {
	db *mongo.Database
}

func NewEmployeeRepository(db *mongo.Database) interfaces.EmployeeRepository {
	return &EmployeeRepository{
		db,
	}
}

func (e *EmployeeRepository) CreateEmployee(employee *dtos.EmployeeDto) (*dtos.EmployeeDto, error) {
	newEntity, err := e.DtoToEntity(employee)
	if err != nil {
		log.Error().Err(err).Msg("error getting new entity")
		return employee, err
	}

	newEntity.CreatedAt = time.Now()

	result, err := e.db.
		Collection("employee").
		InsertOne(context.TODO(), newEntity)
	if err == nil {
		newEntity.ID = result.InsertedID.(primitive.ObjectID)
	} else {
		log.Error().Err(err).Msg("error creating")
		return employee, err
	}

	employee.ID = newEntity.ID.Hex()

	return employee, nil
}

func (r *EmployeeRepository) DtoToEntity(dto *dtos.EmployeeDto) (entity *entities.CronjobEntity, err error) {
	newEntity := entities.CronjobEntity{
		Name: dto.Name,
	}

	if dto.ID != "" {
		objectId, err := primitive.ObjectIDFromHex(dto.ID)
		if err != nil {
			log.Error().Caller().Err(err).Send()
			return entity, err
		}

		newEntity.ID = objectId
	}

	return &newEntity, nil
}
