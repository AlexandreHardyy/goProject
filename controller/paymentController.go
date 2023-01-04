package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/AlexandreHardyy/goProject/database"
	"github.com/AlexandreHardyy/goProject/models"

	"github.com/AlexandreHardyy/goProject/broadcaster"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// CreatePayment godoc
// @Summary Create a payment
// @Description Create a payment
// @Tags payments
// @Accept  json
// @Produce  json
// @Success 200 {object} Payment
// @Router /payments [post]
func CreatePayment(c *gin.Context) {
	payment := &models.Payment{}
	//body, _ := ioutil.ReadAll(c.Request.Body)
	println(c.Params.ByName("ProductID"))
	c.ShouldBindJSON(payment)
	payment.CreatedAt = time.Now()
	payment.UpdatedAt = time.Now()
	fmt.Println(payment.ProductID)
	payment.ProductID = 1

	database.DB.Create(&payment)

	broadcaster := broadcaster.GetBroadcaster()
	broadcaster.Submit(payment)

	// Return a success message if the payment was successfully inserted
	c.JSON(http.StatusOK, gin.H{
		"message": "Payment created",
	})
}

// UpdatePayment godoc
// @Summary Update a payment
// @Description Update a payment
// @Tags payments
// @Accept  json
// @Produce  json
// @Success 200 {object} Payment
// @Router /payments/{id} [put]
func UpdatePayment(c *gin.Context) {
	payment := &models.Payment{}

	c.BindJSON(payment)
	payment.UpdatedAt = time.Now()

	database.DB.Find(&payment, c.Param("id"))
	database.DB.Save(&payment)

	// Return a success message if the payment was successfully updated
	c.JSON(http.StatusOK, gin.H{
		"message": "Payment updated",
	})
}

// DeletePayment godoc
// @Summary Delete a payment
// @Description Delete a payment
// @Tags payments
// @Accept  json
// @Produce  json
// @Success 200 {object} Payment
// @Router /payments/{id} [delete]
func DeletePayment(c *gin.Context) {
	database.DB.Delete(&models.Payment{}, c.Param("id"))
	c.JSON(http.StatusOK, gin.H{
		"message": "Payment deleted",
	})
}

// GetByIdPayment godoc
// @Summary Get a payment
// @Description Get a payment
// @Tags payments
// @Accept  json
// @Produce  json
// @Success 200 {object} Payment
// @Router /payments/{id} [get]
func GetByIdPayment(c *gin.Context) {
	// Initialize a new payment struct
	var payment models.Payment

	database.DB.Find(&payment, c.Param("id"))

	// Return the payment information
	c.JSON(http.StatusOK, gin.H{
		"message": payment,
	})
}

// GetAllPayment godoc
// @Summary Get all payments
// @Description Get all payments
// @Tags payments
// @Accept  json
// @Produce  json
// @Success 200 {object} Payment
// @Router /payments [get]
func GetAllPayment(c *gin.Context) {
	var paymentArray []models.Payment

	database.DB.Find(&paymentArray)

	// Return the productTab slice as a response
	c.JSON(http.StatusOK, gin.H{
		"payment": paymentArray,
	})
}
