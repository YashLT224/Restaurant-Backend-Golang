package routes

import (
	controller "restaurant/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemsRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/ordersItems", controller.GetOrderItems())
	incomingRoutes.GET("/ordersItems/:orderItem_id", controller.GetOrderItem())
	incomingRoutes.GET("/ordersItems-order/:order_id", controller.GetOrderItemsByOrder())
	incomingRoutes.POST("/ordersItems", controller.CreateOrderItem())
	incomingRoutes.PATCH("/ordersItems/:orderItem_id", controller.UpdateOrderItem())
}
