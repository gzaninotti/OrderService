package mapper

import (
	orderitemservice "asdf/tcc/order-service/pkg/domain/order/entity"
	"asdf/tcc/order-service/pkg/network/http/order/serializers"
)

func ToOrderCreationSerializer(order orderitemservice.Order) serializers.Order {
	return serializers.Order{
		UserID:       order.UserID,
		RestaurantID: order.RestaurantID,
		TableID:      order.TableID,
		Status:       order.Status,
		CreatedAt:    order.CreatedAt,
	}
}

func ToOrderUpdateSerializer(order orderitemservice.Order) serializers.Order {
	return serializers.Order{
		ID:        order.ID,
		Status:    order.Status,
		UpdatedAt: order.UpdatedAt,
	}
}

func ToOrderSerializer(order orderitemservice.Order) serializers.Order {
	return serializers.Order{
		ID:           order.ID,
		UserID:       order.UserID,
		RestaurantID: order.RestaurantID,
		TableID:      order.TableID,
		Status:       order.Status,
		CreatedAt:    order.CreatedAt,
		UpdatedAt:    order.UpdatedAt,
	}
}

func ToManyOrderSerializer(orderItens []orderitemservice.Order) []serializers.Order {
	orderItensSerializer := make([]serializers.Order, 0)
	for _, order := range orderItens {
		orderItensSerializer = append(orderItensSerializer, ToOrderSerializer(order))
	}

	return orderItensSerializer
}
