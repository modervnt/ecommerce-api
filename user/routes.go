package user

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine, handler *Handler) {
	r.GET("/users/:id", handler.GetUser)
	r.POST("/users", handler.CreateUser)
	r.POST("/users/login", handler.LoginUser)
	//Penser a ajouter une fonction getAlluser
}
