package controller

import (
	"database/sql"
	"goProject/database"
	"net/http"

	"goProject/broadcaster"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Payment struct {
	Id        int       `json:"id"`
	ProductId int       `json:"productId"`
	PricePaid float64   `json:"pricePaid"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func CreatePayment(c *gin.Context) {
	// Bind the request body to the Payment struct
	var payment Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		// Return a Bad Request error if the request body is invalid
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert the payment into the database
	_, err := database.DB.Exec(`INSERT INTO payment (ProductId, PricePaid, createdAt, updatedAt) VALUES ($1, $2, now(), now())`, payment.ProductId, payment.PricePaid)
	if err != nil {
		// Return an Internal Server Error if there was a problem inserting the payment
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	broadcaster := broadcaster.GetBroadcaster()
	broadcaster.Submit(payment)

	// Return a success message if the payment was successfully inserted
	c.JSON(http.StatusOK, gin.H{
		"message": "Payment created",
	})
}

func UpdatePayment(c *gin.Context) {
	// Bind the request body to the Payment struct
	var payment Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		// Return a Bad Request error if the request body is invalid
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the payment in the database
	_, err := database.DB.Exec(`UPDATE payment SET ProductId=$1, PricePaid=$2, updatedAt=now() WHERE id=$3`, payment.ProductId, payment.PricePaid, c.Param("id"))
	if err != nil {
		// Return an Internal Server Error if there was a problem updating the payment
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success message if the payment was successfully updated
	c.JSON(http.StatusOK, gin.H{
		"message": "Payment updated",
	})
}

func DeletePayment(c *gin.Context) {
	// Delete the payment from the database
	_, err := database.DB.Exec(`DELETE FROM payment where id=$1`, c.Param("id"))
	if err != nil {
		// Return an Internal Server Error if there was a problem deleting the payment
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return a success message if the payment was successfully deleted
	c.JSON(http.StatusOK, gin.H{
		"message": "Payment deleted",
	})
}

func GetByIdPayment(c *gin.Context) {
	// Initialize a new payment struct
	payment := &Payment{}

	// Query the database for the payment with the specified ID
	err := database.DB.QueryRow(`SELECT * FROM payment where id=$1`, c.Param("id")).Scan(&payment.Id, &payment.ProductId, &payment.PricePaid, &payment.CreatedAt, &payment.UpdatedAt)
	if err != nil {
		// Return a Not Found error if the payment does not exist
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
			return
		}
		CheckError(err)
	}

	// Return the payment information
	c.JSON(http.StatusOK, gin.H{
		"message": payment,
	})
}

func GetAllPayment(c *gin.Context) {
	var payment Payment
	var paymentArray []Payment

	// Get all payments from the database
	rows, err := database.DB.Query(`SELECT * FROM payment`)
	if err != nil {
		CheckError(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	// Iterate through the rows and append the payments to the productTab slice
	for rows.Next() {
		err = rows.Scan(&payment.Id, &payment.ProductId, &payment.PricePaid, &payment.CreatedAt, &payment.UpdatedAt)
		if err != nil {
			CheckError(err)
		}
		paymentArray = append(paymentArray, payment)
	}

	// Check for any errors during the iteration
	if err = rows.Err(); err != nil {
		CheckError(err)
	}

	// Return the productTab slice as a response
	c.JSON(http.StatusOK, gin.H{
		"payment": paymentArray,
	})
}
