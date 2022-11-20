package entity

import (
	"github.com/google/uuid"
)

type orderItem struct {
	ID           uuid.UUID `json:"user_id"`
	RestaurantID uuid.UUID `json:"restaurant_id"`
	OrderID      uuid.UUID `json:"order_id"`
	MenuItemID   uuid.UUID `json:"menu_item_id"`
	Quantity     int       `json:"quantity"`
	Price        float64   `json:"price"`
	Status       int       `json:"status"`
}
