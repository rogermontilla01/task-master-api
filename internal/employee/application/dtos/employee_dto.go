package dtos

import (
	"time"

	"github.com/rs/zerolog/log"
)

type EmployeeDto struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Skills         []string  `json:"skills"`
	AvailableHours string    `json:"availableHours"`
	CreatedAt      time.Time `json:"createdAt,omitempty"`
	UpdatedAt      time.Time `json:"updatedAt,omitempty"`
	DeletedAt      time.Time `json:"deletedAt,omitempty"`
}

func (e *EmployeeDto) GetAvailableHours() (duration time.Duration, err error) {
	availableTime, err := time.ParseDuration(e.AvailableHours)
	if err != nil {
		log.Error().Err(err).Msg("error parsing available time")
		return duration, err
	}

	return availableTime, nil
}
