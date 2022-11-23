package mapper

import (
	orderitemservice "asdf/tcc/order-service/pkg/domain/order-item/entity"
	"asdf/tcc/order-service/pkg/network/http/order-item/serializers"
)

func ToOrderItemCreationSerializer(orderItem orderitemservice.OrderItem) serializers.OrderItem {
	return serializers.OrderItem{
		Name:     orderItem.Name,
		Quantity: orderItem.Quantity,
		Price:    orderItem.Price,
		Status:   orderItem.Status,
	}
}

func ToOrderItemUpdateSerializer(orderItem orderitemservice.OrderItem) serializers.OrderItem {
	return serializers.OrderItem{
		Name:   orderItem.Name,
		Status: orderItem.Status,
	}
}

func ToOrderItemSerializer(orderItem orderitemservice.OrderItem) serializers.OrderItem {
	return serializers.OrderItem{
		ID:       orderItem.ID,
		OrderID:  orderItem.OrderID,
		Name:     orderItem.Name,
		Quantity: orderItem.Quantity,
		Price:    orderItem.Price,
		Status:   orderItem.Status,
	}
}

func ToManyOrderItemSerializer(orderItens []orderitemservice.OrderItem) []serializers.OrderItem {
	orderItensSerializer := make([]serializers.OrderItem, 0)
	for _, orderItem := range orderItens {
		orderItensSerializer = append(orderItensSerializer, ToOrderItemSerializer(orderItem))
	}

	return orderItensSerializer
}
