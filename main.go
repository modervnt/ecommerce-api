package main

import (
	"ecommerce-api/db"
	"ecommerce-api/user"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	userStore := user.NewStore(db.DB)

	r := gin.Default()

	user.SetupRoutes(r, user.NewHandler(userStore))
	r.Run(":8080")
}
