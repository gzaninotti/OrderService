package handler

import (
	orderitem "asdf/tcc/order-service/pkg/domain/order"
	"asdf/tcc/order-service/pkg/network/http/order/serializers"
	"asdf/tcc/order-service/pkg/network/http/order/serializers/mapper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OrderHandler interface {
	CreateNewOrder() echo.HandlerFunc
	UpdateOrderStatus() echo.HandlerFunc
	GetOrder() echo.HandlerFunc
	ListUserOrders() echo.HandlerFunc
	ListRestaurantOrders() echo.HandlerFunc
}

type orderHandler struct {
	orderService orderitem.OrderService
}

func NewOrderHandler(service orderitem.OrderService) OrderHandler {
	return &orderHandler{
		orderService: service,
	}
}

func (h *orderHandler) CreateNewOrder() echo.HandlerFunc {
	return func(c echo.Context) error {
		var order serializers.OrderCreation

		if err := c.Bind(&order); err != nil {
			return err
		}

		orderReq := mapper.FromOrderCreationSerializer(order)
		orderRes, err := h.orderService.CreateNewOrder(c.Request().Context(), orderReq)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, mapper.ToOrderCreationSerializer(orderRes))
	}
}

func (h *orderHandler) UpdateOrderStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		var order serializers.OrderStatusUpdate

		if err := c.Bind(order); err != nil {
			return err
		}

		orderReq := mapper.FromOrderUpdateSerializer(order)
		orderRes, err := h.orderService.EditOrderStatus(c.Request().Context(), orderReq)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, mapper.ToOrderUpdateSerializer(orderRes))

	}
}

func (h *orderHandler) GetOrder() echo.HandlerFunc {
	return func(c echo.Context) error {
		var order serializers.Order

		if err := c.Bind(order); err != nil {
			return err
		}

		orderReq := mapper.FromOrderSerializer(order)
		orderRes, err := h.orderService.GetOrder(c.Request().Context(), orderReq)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, mapper.ToOrderSerializer(orderRes))
	}
}

func (h *orderHandler) ListUserOrders() echo.HandlerFunc {
	return func(c echo.Context) error {
		var order serializers.Order

		if err := c.Bind(order); err != nil {
			return err
		}

		orderReq := mapper.FromOrderSerializer(order)
		orderRes, err := h.orderService.GetUserOrders(c.Request().Context(), orderReq)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, mapper.ToManyOrderSerializer(orderRes))
	}
}

func (h *orderHandler) ListRestaurantOrders() echo.HandlerFunc {
	return func(c echo.Context) error {
		var order serializers.Order

		if err := c.Bind(order); err != nil {
			return err
		}

		orderReq := mapper.FromOrderSerializer(order)
		orderRes, err := h.orderService.GetRestaurantOrders(c.Request().Context(), orderReq)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, mapper.ToManyOrderSerializer(orderRes))
	}
}
