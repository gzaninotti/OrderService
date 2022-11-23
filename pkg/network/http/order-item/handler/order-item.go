package handler

import (
	orderitem "asdf/tcc/order-service/pkg/domain/order-item"
	"asdf/tcc/order-service/pkg/network/http/order-item/serializers"
	"asdf/tcc/order-service/pkg/network/http/order-item/serializers/mapper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OrderItemHandler interface {
	CreateNewOrderItem() echo.HandlerFunc
	UpdateOrderItemStatus() echo.HandlerFunc
	GetOrderItem() echo.HandlerFunc
	ListOrderItens() echo.HandlerFunc
}

type orderItemHandler struct {
	orderItemService orderitem.OrderItemService
}

func NewOrderItemHandler(service orderitem.OrderItemService) OrderItemHandler {
	return &orderItemHandler{
		orderItemService: service,
	}
}

func (h *orderItemHandler) CreateNewOrderItem() echo.HandlerFunc {
	return func(c echo.Context) error {
		var orderItem serializers.OrderItemCreation

		if err := c.Bind(&orderItem); err != nil {
			return err
		}

		orderItemReq := mapper.FromOrderItemCreationSerializer(orderItem)
		orderItemRes, err := h.orderItemService.CreateNewOrderItem(c.Request().Context(), orderItemReq)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, mapper.ToOrderItemCreationSerializer(orderItemRes))
	}
}

func (h *orderItemHandler) UpdateOrderItemStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		var orderItem serializers.OrderItemUpdate

		if err := c.Bind(orderItem); err != nil {
			return err
		}

		orderItemReq := mapper.FromOrderItemUpdateSerializer(orderItem)
		orderItemRes, err := h.orderItemService.EditOrderItemStatus(c.Request().Context(), orderItemReq)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, mapper.ToOrderItemUpdateSerializer(orderItemRes))

	}
}

func (h *orderItemHandler) GetOrderItem() echo.HandlerFunc {
	return func(c echo.Context) error {
		var orderItem serializers.OrderItem

		if err := c.Bind(orderItem); err != nil {
			return err
		}

		orderItemReq := mapper.FromOrderItemSerializer(orderItem)
		orderItemRes, err := h.orderItemService.GetOrderItem(c.Request().Context(), orderItemReq)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, mapper.ToOrderItemSerializer(orderItemRes))
	}
}

func (h *orderItemHandler) ListOrderItens() echo.HandlerFunc {
	return func(c echo.Context) error {
		var orderItem serializers.OrderItem

		if err := c.Bind(orderItem); err != nil {
			return err
		}

		orderItemReq := mapper.FromOrderItemSerializer(orderItem)
		orderItemRes, err := h.orderItemService.ListOrderItens(c.Request().Context(), orderItemReq)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, mapper.ToManyOrderItemSerializer(orderItemRes))
	}
}
