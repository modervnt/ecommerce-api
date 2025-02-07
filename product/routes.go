package product

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine, handler *Handler) {
	//r.GET("/products/:id", handler.GetProduct)
	r.POST("/product", handler.CreateProduct)
	//r.DELETE("/product/:id", handler.DeleteByID)
	//r.PATCH("/product/:id", handler.UpdateProduct)
}
