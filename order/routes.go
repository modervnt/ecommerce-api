package order

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine, handler *Handler) {
	r.GET("/order/:id", handler.GetOrderByID)
	r.POST("/order", handler.CreateOrder)
	r.DELETE("/order/:id", handler.DeleteOrderByID)
	r.PATCH("/order/:id", handler.UpdateOrder)
}
