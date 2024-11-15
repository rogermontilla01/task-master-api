package repository

import (
	"context"
	"task-master-api/internal/assignment/application/dtos"
	"task-master-api/internal/assignment/domain/interfaces"
	"task-master-api/internal/assignment/infrastructure/entities"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AssignmentRepository struct {
	db *mongo.Database
}

func NewAssignmentRepository(db *mongo.Database) interfaces.AssignmentRepository {
	return &AssignmentRepository{db}
}

func (a *AssignmentRepository) CreateAssignment(assignment *dtos.AssignmentDto) (*dtos.AssignmentDto, error) {
	newEntity, err := a.DtoToEntity(assignment)
	if err != nil {
		log.Error().Err(err).Msg("error converting dto to entity")
		return nil, err
	}

	newEntity.CreatedAt = time.Now()

	result, err := a.db.
		Collection("assignments").
		InsertOne(context.TODO(), newEntity)
	if err == nil {
		newEntity.ID = result.InsertedID.(primitive.ObjectID)
	} else {
		log.Error().Err(err).Msg("error inserting assignment")
		return nil, err
	}

	assignment.ID = newEntity.ID.Hex()
	assignment.CreatedAt = newEntity.CreatedAt

	return assignment, nil
}

func (a *AssignmentRepository) DeleteAssignment(id string) (err error) {
	var objectId primitive.ObjectID

	if id != "" {
		objectId, err = primitive.ObjectIDFromHex(id)

		if err != nil {
			log.Error().Caller().Err(err).Send()
			return err
		}
	}

	_, err = a.db.
		Collection("assignments").
		DeleteOne(context.TODO(), bson.M{"_id": objectId})
	if err != nil {
		log.Error().Caller().Err(err).Send()
		return err
	}

	return nil
}

func (a *AssignmentRepository) GetAllAssignments() (*[]dtos.AssignmentDto, error) {
	entities := []entities.AssignmentEntity{}

	cursor, err := a.db.
		Collection("assignments").
		Find(context.TODO(), bson.M{})
	if err != nil {
		log.Error().Caller().Err(err).Send()
		return nil, err
	}

	err = cursor.All(context.TODO(), &entities)
	if err != nil {
		log.Error().Caller().Err(err).Send()
		return nil, err
	}

	allAssignments := []dtos.AssignmentDto{}
	for _, entity := range entities {
		assignmentDto, err := a.EntityToDto(&entity)
		if err != nil {
			log.Error().Caller().Err(err).Send()
			return nil, err
		}

		allAssignments = append(allAssignments, *assignmentDto)
	}

	return &allAssignments, nil
}

func (a *AssignmentRepository) GetAssignmentById(id string) (*dtos.AssignmentDto, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Error().Caller().Err(err).Send()
		return nil, err
	}

	entity := entities.AssignmentEntity{}
	err = a.db.
		Collection("assignments").
		FindOne(context.TODO(), bson.M{"_id": objectId}).
		Decode(&entity)
	if err != nil {
		log.Error().Caller().Err(err).Send()
		return nil, err
	}

	dto, err := a.EntityToDto(&entity)
	if err != nil {
		log.Error().Caller().Err(err).Send()
		return nil, err
	}

	return dto, nil
}

func (a *AssignmentRepository) UpdateAssignment(id string, assignment *dtos.UpdateAssignmentDto) (updated *dtos.UpdateAssignmentDto, err error) {
	var objectId primitive.ObjectID

	if id != "" {
		objectId, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			log.Error().Caller().Err(err).Send()
			return nil, err
		}
	}

	update := bson.M{}

	if assignment.Duration != nil {
		update["duration"] = assignment.Duration
	}

	if assignment.EmployeeID != nil {
		objectId, err = primitive.ObjectIDFromHex(*assignment.EmployeeID)
		if err != nil {
			log.Error().Caller().Err(err).Send()
			return nil, err
		}

		update["employeeId"] = objectId
	}

	if assignment.TaskID != nil {
		objectId, err = primitive.ObjectIDFromHex(*assignment.TaskID)
		if err != nil {
			log.Error().Caller().Err(err).Send()
			return nil, err
		}

		update["taskId"] = objectId
	}

	update["updatedAt"] = time.Now()

	_, err = a.db.
		Collection("assignments").
		UpdateOne(context.TODO(), bson.M{"_id": objectId}, bson.M{"$set": update})
	if err != nil {
		log.Error().Caller().Err(err).Send()
		return nil, err
	}

	return assignment, nil
}

func (a *AssignmentRepository) DtoToEntity(dto *dtos.AssignmentDto) (*entities.AssignmentEntity, error) {
	assignmentEntity := entities.AssignmentEntity{
		Duration: dto.Duration,
	}

	if dto.ID != "" {
		objectId, err := primitive.ObjectIDFromHex(dto.ID)
		if err != nil {
			log.Error().Caller().Err(err).Send()
			return nil, err
		}

		assignmentEntity.ID = objectId
	}

	if dto.EmployeeID != "" {
		objectId, err := primitive.ObjectIDFromHex(dto.ID)
		if err != nil {
			log.Error().Caller().Err(err).Send()
			return nil, err
		}

		assignmentEntity.EmployeeID = objectId
	}

	if dto.TaskID != "" {
		objectId, err := primitive.ObjectIDFromHex(dto.ID)
		if err != nil {
			log.Error().Caller().Err(err).Send()
			return nil, err
		}

		assignmentEntity.TaskID = objectId
	}

	return &assignmentEntity, nil
}

func (a *AssignmentRepository) EntityToDto(entity *entities.AssignmentEntity) (*dtos.AssignmentDto, error) {
	assignmentDto := dtos.AssignmentDto{
		ID:         entity.ID.Hex(),
		EmployeeID: entity.EmployeeID.Hex(),
		TaskID:     entity.TaskID.Hex(),
		Duration:   entity.Duration,
	}

	return &assignmentDto, nil
}
