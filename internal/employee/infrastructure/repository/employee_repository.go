package repository

import (
	"context"
	"task-master-api/internal/employee/application/dtos"
	"task-master-api/internal/employee/domain/interfaces"
	"task-master-api/internal/employee/infrastructure/entities"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeRepository struct {
	db *mongo.Database
}

func NewEmployeeRepository(db *mongo.Database) interfaces.EmployeeRepository {
	return &EmployeeRepository{db}
}

func (e *EmployeeRepository) GetEmployee(id string) (*dtos.EmployeeDto, error) {
	employee := entities.EmployeeEntity{}

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Error().Caller().Err(err).Send()
		return nil, err
	}

	err = e.db.
		Collection("employee").
		FindOne(context.TODO(), bson.M{"_id": objectId}).
		Decode(&employee)
	if err != nil {
		log.Error().Caller().Err(err).Send()
		return nil, err
	}

	dto, err := e.EntityToDto(&employee)
	if err != nil {
		log.Error().Caller().Err(err).Send()
		return nil, err
	}

	return dto, nil
}

func (e *EmployeeRepository) CreateEmployee(employee *dtos.EmployeeDto) (*dtos.EmployeeDto, error) {
	newEntity, err := e.DtoToEntity(employee)
	if err != nil {
		log.Error().Err(err).Msg("error getting new entity")
		return nil, err
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
	employee.CreatedAt = newEntity.CreatedAt

	return employee, nil
}

func (e *EmployeeRepository) UpdateEmployee(id string, employee *dtos.UpdateEmployeeDto) (updated *dtos.UpdateEmployeeDto, err error) {
	var objectId primitive.ObjectID

	if id != "" {
		objectId, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			log.Error().Caller().Err(err).Send()
			return nil, err
		}
	}

	update := bson.M{}

	if employee.Name != nil {
		update["name"] = employee.Name
	}
	if employee.Skills != nil {
		update["skills"] = employee.Skills
	}
	if employee.AvailableHours != nil {
		update["availableHours"] = employee.AvailableHours
	}
	if employee.AvailableDays != nil {
		update["availableDays"] = employee.AvailableDays
	}
	update["updatedAt"] = time.Now()

	_, err = e.db.
		Collection("employee").
		UpdateOne(context.TODO(), bson.M{"_id": objectId}, bson.M{"$set": update})
	if err != nil {
		log.Error().Caller().Err(err).Send()
		return nil, err
	}

	return employee, nil
}

func (r *EmployeeRepository) DeleteEmployee(id string) (err error) {
	var objectId primitive.ObjectID

	if id != "" {
		objectId, err = primitive.ObjectIDFromHex(id)

		if err != nil {
			log.Error().Caller().Err(err).Send()
			return err
		}
	}

	_, err = r.db.
		Collection("employee").
		DeleteOne(context.TODO(), bson.M{"_id": objectId})
	if err != nil {
		log.Error().Caller().Err(err).Send()
		return err
	}

	return nil
}

func (r *EmployeeRepository) DtoToEntity(dto *dtos.EmployeeDto) (entity *entities.EmployeeEntity, err error) {
	newEntity := entities.EmployeeEntity{
		Name:           dto.Name,
		Skills:         dto.Skills,
		AvailableHours: dto.AvailableHours,
		AvailableDays:  dto.AvailableDays,
	}

	if dto.ID != "" {
		objectId, err := primitive.ObjectIDFromHex(dto.ID)
		if err != nil {
			log.Error().Caller().Err(err).Send()
			return nil, err
		}

		newEntity.ID = objectId
	}

	return &newEntity, nil
}

func (r *EmployeeRepository) EntityToDto(entity *entities.EmployeeEntity) (dto *dtos.EmployeeDto, err error) {
	newDto := dtos.EmployeeDto{
		ID:             entity.ID.Hex(),
		Name:           entity.Name,
		Skills:         entity.Skills,
		AvailableHours: entity.AvailableHours,
		AvailableDays:  entity.AvailableDays,
		CreatedAt:      entity.CreatedAt,
		UpdatedAt:      entity.UpdatedAt,
		DeletedAt:      entity.DeletedAt,
	}

	return &newDto, nil
}
