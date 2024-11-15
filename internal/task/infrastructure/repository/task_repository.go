package repository

import (
	"context"
	"task-master-api/internal/task/application/dtos"
	"task-master-api/internal/task/domain/interfaces"
	"task-master-api/internal/task/infrastructure/entities"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository struct {
	db *mongo.Database
}

func NewTaskRepository(db *mongo.Database) interfaces.TaskRepository {
	return &TaskRepository{db}
}

func (t *TaskRepository) CreateTask(task *dtos.TaskDto) (*dtos.TaskDto, error) {
	newEntity, err := t.DtoToEntity(task)
	if err != nil {
		log.Error().Err(err).Msg("Error converting dto to entity")
		return nil, err
	}

	newEntity.CreatedAt = time.Now()

	result, err := t.db.
		Collection("tasks").
		InsertOne(context.TODO(), newEntity)
	if err == nil {
		newEntity.ID = result.InsertedID.(primitive.ObjectID)
	} else {
		log.Error().Err(err).Msg("error creating")
		return task, err
	}

	task.ID = newEntity.ID.Hex()
	task.CreatedAt = newEntity.CreatedAt

	return task, nil
}

func (t *TaskRepository) DeleteTask(id string) (err error) {
	var objectId primitive.ObjectID

	if id != "" {
		objectId, err = primitive.ObjectIDFromHex(id)

		if err != nil {
			log.Error().Caller().Err(err).Send()
			return err
		}
	}

	_, err = t.db.
		Collection("task").
		DeleteOne(context.TODO(), bson.M{"_id": objectId})
	if err != nil {
		log.Error().Caller().Err(err).Send()
		return err
	}

	return nil
}

func (t *TaskRepository) GetAllTasks() (*[]dtos.TaskDto, error) {
	entities := []entities.TaskEntity{}

	cursor, err := t.db.
		Collection("tasks").
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

	allTasks := []dtos.TaskDto{}
	for _, entity := range entities {
		dto, err := t.EntityToDto(&entity)
		if err != nil {
			log.Error().Caller().Err(err).Send()
			return nil, err
		}

		allTasks = append(allTasks, *dto)
	}

	return &allTasks, nil
}

func (t *TaskRepository) GetTaskById(id string) (*dtos.TaskDto, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Error().Caller().Err(err).Send()
		return nil, err
	}

	entity := entities.TaskEntity{}
	err = t.db.
		Collection("tasks").
		FindOne(context.TODO(), bson.M{"_id": objectId}).
		Decode(&entity)
	if err != nil {
		log.Error().Caller().Err(err).Send()
		return nil, err
	}

	dto, err := t.EntityToDto(&entity)
	if err != nil {
		log.Error().Caller().Err(err).Send()
		return nil, err
	}

	return dto, nil
}

func (t *TaskRepository) UpdateTask(id string, task *dtos.UpdateTaskDto) (updated *dtos.UpdateTaskDto, err error) {
	var objectId primitive.ObjectID

	if id != "" {
		objectId, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			log.Error().Caller().Err(err).Send()
			return nil, err
		}
	}

	update := bson.M{}

	if task.Title != nil {
		update["title"] = task.Title
	}

	if task.Duration != nil {
		update["duration"] = task.Duration
	}

	if task.Skills != nil {
		update["skills"] = task.Skills
	}

	if task.Completed != nil {
		update["completed"] = task.Completed
	}

	update["updatedAt"] = time.Now()

	_, err = t.db.
		Collection("tasks").
		UpdateOne(context.TODO(), bson.M{"_id": objectId}, bson.M{"$set": update})
	if err != nil {
		log.Error().Caller().Err(err).Send()
		return nil, err
	}

	return task, nil
}

func (t *TaskRepository) DtoToEntity(dto *dtos.TaskDto) (*entities.TaskEntity, error) {
	newEntity := entities.TaskEntity{
		Title:     dto.Title,
		Duration:  dto.Duration,
		Skills:    dto.Skills,
		Completed: dto.Completed,
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

func (t *TaskRepository) EntityToDto(entity *entities.TaskEntity) (*dtos.TaskDto, error) {
	newDto := dtos.TaskDto{
		ID:        entity.ID.Hex(),
		Title:     entity.Title,
		Duration:  entity.Duration,
		Skills:    entity.Skills,
		Completed: entity.Completed,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		DeletedAt: entity.DeletedAt,
	}

	return &newDto, nil
}
