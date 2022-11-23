package mapper

import (
	orderitemservice "asdf/tcc/order-service/pkg/domain/order/entity"
	"asdf/tcc/order-service/pkg/network/http/order/serializers"
)

func FromOrderCreationSerializer(order serializers.OrderCreation) orderitemservice.Order {
	return orderitemservice.Order{
		UserID:       order.UserID,
		RestaurantID: order.RestaurantID,
		TableID:      order.TableID,
	}
}

func FromOrderUpdateSerializer(order serializers.OrderStatusUpdate) orderitemservice.Order {
	return orderitemservice.Order{
		ID:           order.ID,
		UserID:       order.UserID,
		RestaurantID: order.RestaurantID,
		TableID:      order.TableID,
		Status:       order.Status,
	}
}

func FromOrderSerializer(order serializers.Order) orderitemservice.Order {
	return orderitemservice.Order{
		ID:           order.ID,
		UserID:       order.UserID,
		RestaurantID: order.RestaurantID,
		TableID:      order.TableID,
	}
}
