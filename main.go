package main

import (
	"goProject/controller"
	"goProject/database"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()

	database.ConnectDatabase()

	// Routes pour les produits
	productGroup := r.Group("/product")
	{
		productGroup.GET("/", controller.GetAllProduct)
		productGroup.GET("/:id", controller.GetByIdProduct)
		productGroup.DELETE("/delete/:id", controller.DeleteProduct)
		productGroup.POST("/create", controller.CreateProduct)
		productGroup.PUT("/update/:id", controller.UpdateProduct)
	}

	// Payment routes
	paymentRoutes := r.Group("/payment")
	{
		paymentRoutes.POST("/create", controller.CreatePayment)
		paymentRoutes.PUT("/update/:id", controller.UpdatePayment)
		paymentRoutes.DELETE("/delete/:id", controller.DeletePayment)
		paymentRoutes.GET("/:id", controller.GetByIdPayment)
		paymentRoutes.GET("/", controller.GetAllPayment)
	}

	r.GET("/stream/payment", controller.Stream)

	r.Run(":3000")
}
