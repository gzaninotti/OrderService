package serializers

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID           uuid.UUID `json:"id,omitempty" bson:"id,omitempty"`
	UserID       uuid.UUID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	RestaurantID uuid.UUID `json:"restaurant_id,omitempty" bson:"restaurant_id,omitempty"`
	TableID      int       `json:"table_id,omitempty" bson:"table_id,omitempty"`
	Status       int       `json:"status,omitempty" bson:"status,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type OrderCreation struct {
	UserID       uuid.UUID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	TableID      int       `json:"table_id,omitempty" bson:"table_id,omitempty"`
	RestaurantID uuid.UUID `json:"restaurant_id,omitempty" bson:"restaurant_id,omitempty"`
}

type OrderStatusUpdate struct {
	ID           uuid.UUID `json:"order_id,omitempty" bson:"order_id,omitempty"`
	UserID       uuid.UUID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	TableID      int       `json:"table_id,omitempty" bson:"table_id,omitempty"`
	RestaurantID uuid.UUID `json:"restaurant_id,omitempty" bson:"restaurant_id,omitempty"`
	Status       int       `json:"status,omitempty" bson:"status,omitempty"`
}
