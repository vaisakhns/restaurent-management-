package routes

import (
	"github.com/gin-gonic/gin"
	controller "golang-restaurent-management/controllers"

)




func OrderItemRoutes(incomingRoutes *gin.Engine)  {
	
	incomingRoutes.GET("/orderItems",controller.GetOrderItems())
	incomingRoutes.GET("/orderItems/:orderItem_id",controller.GetOrderItem())
	incomingRoutes.GET("/orderItems-order/:order_id",controller.GetOrderItemsByOrder())
    incomingRoutes.POST("/orderItems",controller.CreateOrderItems())
	incomingRoutes.PATCH("/orderItems/:orderItems_id",controller. UpdateOrderItems())
	
}