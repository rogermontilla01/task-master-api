package dtos

import (
	"time"

	"github.com/rs/zerolog/log"
)

type TaskDto struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Duration  string    `json:"duration,omitempty"`
	Skills    []string  `json:"skills"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
}

func (t *TaskDto) GetDuration() (duration time.Duration, err error) {
	availableTime, err := time.ParseDuration(t.Duration)
	if err != nil {
		log.Error().Err(err).Msg("error parsing duration time")
		return duration, err
	}

	return availableTime, nil
}
