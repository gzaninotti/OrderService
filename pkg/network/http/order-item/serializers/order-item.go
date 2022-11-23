package serializers

import "github.com/google/uuid"

type OrderItem struct {
	ID       uuid.UUID `json:"id,omitempty" bson:"id,omitempty"`
	OrderID  uuid.UUID `json:"order_id,omitempty" bson:"order_id,omitempty"`
	Name     string    `json:"name,omitempty" bson:"name,omitempty"`
	Quantity int       `json:"quantity,omitempty" bson:"quantity,omitempty"`
	Price    float64   `json:"price,omitempty" bson:"price,omitempty"`
	Status   int       `json:"status,omitempty" bson:"status,omitempty"`
}

type OrderItemCreation struct {
	OrderID   uuid.UUID `json:"order_id,omitempty" bson:"order_id,omitempty"`
	Name      string    `json:"name,omitempty" bson:"name,omitempty"`
	Quantity  int       `json:"quantity,omitempty" bson:"quantity,omitempty"`
	UnitPrice float64   `json:"unit_price,omitempty" bson:"unit_price,omitempty"`
}

type OrderItemUpdate struct {
	OrderID uuid.UUID `json:"order_id,omitempty" bson:"order_id,omitempty"`
	Status  int       `json:"status,omitempty" bson:"status,omitempty"`
}
