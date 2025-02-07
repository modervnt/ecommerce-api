package user

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine, handler *Handler) {
	r.GET("/users/:id", handler.GetUser)
	r.POST("/users", handler.CreateUser)
}
