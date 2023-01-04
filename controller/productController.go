package controller

import (
	"database/sql"
	"goProject/database"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Product struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Price     float32 `json:"price"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
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
	_ = c.BindJSON(prod)
	_, err := database.DB.Exec(`INSERT INTO product (name, price, createdAt, updatedAt) VALUES ($1, $2, now(), now())`, prod.Name, prod.Price)
	CheckError(err)
	c.JSON(http.StatusOK, gin.H{
		"message": "Product created",
	})
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
	c.BindJSON(prod)
	_, err := database.DB.Exec(`UPDATE product SET name=$1, price=$2, updatedAt=now() WHERE id=$3`, prod.Name, prod.Price, c.Param("id"))
	CheckError(err)
	c.JSON(http.StatusOK, gin.H{
		"message": "Product updated",
	})
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
