package entity

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	RestaurantID uuid.UUID
	Status       int
	TableID      int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
