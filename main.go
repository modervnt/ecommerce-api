package main

import (
	"ecommerce-api/db"
	"ecommerce-api/order"
	"ecommerce-api/product"
	"ecommerce-api/user"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	userStore := user.NewStore(db.DB)
	productStore := product.NewStore(db.DB)
	orderStore := order.NewStore(db.DB)

	r := gin.Default()

	user.SetupRoutes(r, user.NewHandler(userStore))
	product.SetupRoutes(r, product.NewHandler(productStore))
	order.SetupRoutes(r, order.NewHandler(orderStore))
	r.Run(":8080")
}
