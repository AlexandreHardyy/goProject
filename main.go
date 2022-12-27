package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "example"
	dbname   = "postgres"
)

type product struct {
	id        int
	name      string
	price     int
	createdAt string
	updatedAt string
}

type Payment struct {
	Id        int       `json:"id"`
	ProductId int       `json:"productId"`
	PricePaid float64   `json:"pricePaid"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func main() {
	r := gin.Default()

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	//psqlconn := "postgres://postgres:example@db/postgres?sslmode=disable"

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// check db
	err = db.Ping()
	CheckError(err)

	product := r.Group("/product")
	{
		product.GET("/", func(c *gin.Context) {
			var res string
			var productTab []string

			rows, err := db.Query(`SELECT name FROM product`)
			CheckError(err)

			defer rows.Close()
			for rows.Next() {
				err = rows.Scan(&res)
				CheckError(err)
				productTab = append(productTab, res)
			}

			c.JSON(http.StatusOK, gin.H{
				"product": productTab,
			})
		})

		product.GET("/:id", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "getBiId",
			})
		})
	}

	// Payment routes
	paymentRoutes := r.Group("/payment")
	{
		// Create
		paymentRoutes.POST("/create", func(c *gin.Context) {
			// Bind the request body to the Payment struct
			var payment Payment
			if err := c.ShouldBindJSON(&payment); err != nil {
				// Return a Bad Request error if the request body is invalid
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			// Insert the payment into the database
			_, err := db.Exec(`INSERT INTO payment (ProductId, PricePaid, createdAt, updatedAt) VALUES ($1, $2, now(), now())`, payment.ProductId, payment.PricePaid)
			if err != nil {
				// Return an Internal Server Error if there was a problem inserting the payment
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			// Return a success message if the payment was successfully inserted
			c.JSON(http.StatusOK, gin.H{
				"message": "Payment created",
			})
		})

		// Update
		paymentRoutes.PUT("/update/:id", func(c *gin.Context) {
			// Bind the request body to the Payment struct
			var payment Payment
			if err := c.ShouldBindJSON(&payment); err != nil {
				// Return a Bad Request error if the request body is invalid
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			// Update the payment in the database
			_, err := db.Exec(`UPDATE payment SET ProductId=$1, PricePaid=$2, updatedAt=now() WHERE id=$3`, payment.ProductId, payment.PricePaid, c.Param("id"))
			if err != nil {
				// Return an Internal Server Error if there was a problem updating the payment
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			// Return a success message if the payment was successfully updated
			c.JSON(http.StatusOK, gin.H{
				"message": "Payment updated",
			})
		})

		// Delete
		paymentRoutes.DELETE("/delete/:id", func(c *gin.Context) {
			// Delete the payment from the database
			_, err := db.Exec(`DELETE FROM payment where id=$1`, c.Param("id"))
			if err != nil {
				// Return an Internal Server Error if there was a problem deleting the payment
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			// Return a success message if the payment was successfully deleted
			c.JSON(http.StatusOK, gin.H{
				"message": "Payment deleted",
			})
		})

		// GetById
		paymentRoutes.GET("/:id", func(c *gin.Context) {
			// Initialize a new payment struct
			payment := &Payment{}

			// Query the database for the payment with the specified ID
			err := db.QueryRow(`SELECT * FROM payment where id=$1`, c.Param("id")).Scan(&payment.Id, &payment.ProductId, &payment.PricePaid, &payment.CreatedAt, &payment.UpdatedAt)
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
		})

		// GetAll
		paymentRoutes.GET("/", func(c *gin.Context) {
			var payment Payment
			var paymentArray []Payment

			// Get all payments from the database
			rows, err := db.Query(`SELECT * FROM payment`)
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
		})
	}

	err = r.Run(":3000")
	if err != nil {
		return
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
