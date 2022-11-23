package route

import (
	"asdf/tcc/order-service/pkg/network/http/order/handler"

	"github.com/labstack/echo/v4"
)

func Configure(e *echo.Echo, orderHandler handler.OrderHandler) {
	ConfigureOrderEndpoints(e, orderHandler)
}

func ConfigureOrderEndpoints(e *echo.Echo, orderHandler handler.OrderHandler) {
	// CREATE Order ENDPOINT
	e.POST("/order/create", orderHandler.CreateNewOrder())
	// UPDATE Order Status ENDPOINT
	e.PUT("/order/update/status", orderHandler.UpdateOrderStatus())
	// GET Order ENDPOINT
	e.GET("/order/get", orderHandler.GetOrder())
	// LIST User Orders ENDPOINT
	e.GET("/orders/user", orderHandler.ListUserOrders())
	// LIST Restaurant Orders ENDPOINT
	e.GET("/orders/restaurant", orderHandler.ListRestaurantOrders())
}
