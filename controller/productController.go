package controller

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/AlexandreHardyy/goProject/database"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Product struct {
	Id        int       `json:"id"`
	Name      string    `json:"name" binding:"required"`
	Price     float32   `json:"price" binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// GetAllProduct godoc
// @Summary Get all products
// @Description Get all products
// @Tags products
// @Accept  json
// @Produce  json
// @Success 200 {object} Product
// @Router /products [get]
func GetAllProduct(c *gin.Context) {
	var prod Product
	var productTab []Product

	rows, err := database.DB.Query(`SELECT * FROM product`)
	CheckError(err)

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)
	for rows.Next() {
		err = rows.Scan(&prod.Id, &prod.Name, &prod.Price, &prod.CreatedAt, &prod.UpdatedAt)
		CheckError(err)
		productTab = append(productTab, prod)
	}

	c.JSON(http.StatusOK, gin.H{
		"product": productTab,
	})
}

// GetByIdProduct godoc
// @Summary Get products by id
// @Description Get products by id
// @Tags products
// @Accept  json
// @Produce  json
// @Success 200 {object} Product
// @Router /products/{id} [get]
func GetByIdProduct(c *gin.Context) {
	var prod Product

	rows, err := database.DB.Query(`SELECT * FROM product where id=$1`, c.Param("id"))
	CheckError(err)

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)
	if rows.Next() {
		err = rows.Scan(&prod.Id, &prod.Name, &prod.Price, &prod.CreatedAt, &prod.UpdatedAt)
		CheckError(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": prod,
	})
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Delete a product
// @Tags products
// @Accept  json
// @Produce  json
// @Success 200 {object} Product
// @Router /products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	_, err := database.DB.Exec(`DELETE FROM product where id=$1`, c.Param("id"))
	CheckError(err)
	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted",
	})
}

// CreateProduct godoc
// @Summary Create a product
// @Description Create a product
// @Tags products
// @Accept  json
// @Produce  json
// @Success 200 {object} Product
// @Router /products [post]
func CreateProduct(c *gin.Context) {
	prod := &Product{}

	prod.CreatedAt = time.Now()
	prod.UpdatedAt = time.Now()
	if err := c.BindJSON(prod); err == nil {
		err := database.DB.QueryRow(`INSERT INTO product (name, price, createdAt, updatedAt) VALUES ($1, $2, $3, $4) RETURNING id`, prod.Name, prod.Price, prod.CreatedAt, prod.UpdatedAt).Scan(&prod.Id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "Product created",
				"product": prod,
			})
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

// UpdateProduct godoc
// @Summary Update a product
// @Description Update a product
// @Tags products
// @Accept  json
// @Produce  json
// @Success 200 {object} Product
// @Router /products/{id} [put]
func UpdateProduct(c *gin.Context) {
	prod := &Product{}
	prod.UpdatedAt = time.Now()

	if err := c.BindJSON(prod); err == nil {
		err := database.DB.QueryRow(`UPDATE product SET name=$1, price=$2, updatedAt=$3 WHERE id=$4 RETURNING id, createdAt, updatedAt`, prod.Name, prod.Price, prod.UpdatedAt, c.Param("id")).Scan(&prod.Id, &prod.CreatedAt, &prod.UpdatedAt)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "Product updated",
				"product": prod,
			})
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

// CheckError godoc
// @Summary Check error
// @Description Check error
// @Tags products
// @Accept  json
// @Produce  json
// @Success 200 {object} Product
// @Router /products/{id} [put]
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
