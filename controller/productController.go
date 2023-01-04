package controller

import (
	"net/http"
	"time"

	"github.com/AlexandreHardyy/goProject/database"
	"github.com/AlexandreHardyy/goProject/models"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// GetAllProduct godoc
// @Summary Get all products
// @Description Get all products
// @Tags products
// @Accept  json
// @Produce  json
// @Success 200 {object} Product
// @Router /products [get]
func GetAllProduct(c *gin.Context) {
	var productTab []models.Product

	database.DB.Find(&productTab)

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
	var prod models.Product

	database.DB.Find(&prod, c.Param("id"))

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
	database.DB.Delete(&models.Product{}, c.Param("id"))
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
	prod := &models.Product{}
	c.BindJSON(prod)
	prod.CreatedAt = time.Now()
	prod.UpdatedAt = time.Now()

	database.DB.Create(&prod)

	c.JSON(http.StatusOK, gin.H{
		"message": "Product created",
		"product": prod,
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
	prod := &models.Product{}
	database.DB.Find(&prod, c.Param("id"))
	c.BindJSON(prod)
	prod.UpdatedAt = time.Now()

	database.DB.Save(&prod)

	c.JSON(http.StatusOK, gin.H{
		"message": "Product updated",
		"product": prod,
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
