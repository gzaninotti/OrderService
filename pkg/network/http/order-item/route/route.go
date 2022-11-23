package route

import (
	"asdf/tcc/order-service/pkg/network/http/order-item/handler"

	"github.com/labstack/echo/v4"
)

func Configure(e *echo.Echo, orderItemHandler handler.OrderItemHandler) {
	ConfigureOrderItemEndpoints(e, orderItemHandler)
}

func ConfigureOrderItemEndpoints(e *echo.Echo, orderItemHandler handler.OrderItemHandler) {
	// CREATE OrderItem ENDPOINT
	e.POST("/order/item/create", orderItemHandler.CreateNewOrderItem())
	// UPDATE OrderItem ENDPOINT
	e.PUT("/order/item/update/status", orderItemHandler.UpdateOrderItemStatus())
	// GET OrderItem ENDPOINT
	e.GET("/order/item/get", orderItemHandler.GetOrderItem())
	// LIST OrderItem ENDPOINT
	e.GET("/order/itens", orderItemHandler.ListOrderItens())
}
