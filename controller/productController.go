package controller

import (
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

func GetAllProduct(c *gin.Context) {
	var prod Product
	productTab := []Product{}

	rows, err := database.DB.Query(`SELECT * FROM product`)
	CheckError(err)

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&prod.Id, &prod.Name, &prod.Price, &prod.CreatedAt, &prod.UpdatedAt)
		CheckError(err)
		productTab = append(productTab, prod)
	}

	c.JSON(http.StatusOK, gin.H{
		"product": productTab,
	})
}

func GetByIdProduct(c *gin.Context) {
	var prod Product

	rows, err := database.DB.Query(`SELECT * FROM product where id=$1`, c.Param("id"))
	CheckError(err)

	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&prod.Id, &prod.Name, &prod.Price, &prod.CreatedAt, &prod.UpdatedAt)
		CheckError(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": prod,
	})
}

func DeleteProduct(c *gin.Context) {
	_, err := database.DB.Exec(`DELETE FROM product where id=$1`, c.Param("id"))
	CheckError(err)
	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted",
	})
}

func CreateProduct(c *gin.Context) {
	prod := &Product{}
	c.BindJSON(prod)
	_, err := database.DB.Exec(`INSERT INTO product (name, price, createdAt, updatedAt) VALUES ($1, $2, now(), now())`, prod.Name, prod.Price)
	CheckError(err)
	c.JSON(http.StatusOK, gin.H{
		"message": "Product created",
	})
}

func UpdateProduct(c *gin.Context) {
	prod := &Product{}
	c.BindJSON(prod)
	_, err := database.DB.Exec(`UPDATE product SET name=$1, price=$2, updatedAt=now() WHERE id=$3`, prod.Name, prod.Price, c.Param("id"))
	CheckError(err)
	c.JSON(http.StatusOK, gin.H{
		"message": "Product updated",
	})
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
