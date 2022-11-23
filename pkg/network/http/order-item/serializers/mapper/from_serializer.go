package mapper

import (
	orderitemservice "asdf/tcc/order-service/pkg/domain/order-item/entity"
	"asdf/tcc/order-service/pkg/network/http/order-item/serializers"
)

func FromOrderItemCreationSerializer(orderItem serializers.OrderItemCreation) orderitemservice.OrderItem {
	return orderitemservice.OrderItem{
		OrderID:   orderItem.OrderID,
		Name:      orderItem.Name,
		Quantity:  orderItem.Quantity,
		UnitPrice: orderItem.UnitPrice,
	}
}

func FromOrderItemUpdateSerializer(orderItem serializers.OrderItemUpdate) orderitemservice.OrderItem {
	return orderitemservice.OrderItem{
		OrderID: orderItem.OrderID,
		Status:  orderItem.Status,
	}
}

func FromOrderItemSerializer(orderItem serializers.OrderItem) orderitemservice.OrderItem {
	return orderitemservice.OrderItem{
		ID:      orderItem.ID,
		OrderID: orderItem.OrderID,
	}
}
