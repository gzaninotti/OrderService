package entity

import (
	"github.com/google/uuid"
)

type OrderItem struct {
	ID        uuid.UUID
	OrderID   uuid.UUID
	Name      string
	Quantity  int
	Price     float64
	Status    int
	UnitPrice float64
}
