package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"goProject/controller"
	"goProject/database"

	_ "goProject/docs"
)

// @title           API in golang
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /
func main() {
	r := gin.Default()

	database.ConnectDatabase()
	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	/*r.GET("/swagger/*any", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Swagger",
		})
	})*/

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

	err := r.Run(":3000")
	if err != nil {
		return
	}
}
