package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "example"
	dbname   = "postgres"
)

type Product struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Price     float32 `json:"price"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}

func main() {
	r := gin.Default()

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// check db
	err = db.Ping()
	CheckError(err)

	// Routes pour les produits
	productGroup := r.Group("/product")
	{
		// GetAll
		productGroup.GET("/", func(c *gin.Context) {
			productTab := GetAll(db)
			c.JSON(http.StatusOK, gin.H{
				"product": productTab,
			})
		})

		// GetById
		productGroup.GET("/:id", func(c *gin.Context) {
			prod := GetById(db, c.Param("id"))
			c.JSON(http.StatusOK, gin.H{
				"message": prod,
			})
		})

		// Delete
		productGroup.DELETE("/delete/:id", func(c *gin.Context) {
			Delete(db, c.Param("id"))
			c.JSON(http.StatusOK, gin.H{
				"message": "Product deleted",
			})
		})

		// Create
		productGroup.POST("/create", func(c *gin.Context) {
			prod := &Product{}
			c.BindJSON(prod)
			fmt.Println(prod)
			Create(db, prod.Name, prod.Price)
			c.JSON(http.StatusOK, gin.H{
				"message": "Product created",
			})
		})

		// Update
		productGroup.PUT("/update/:id", func(c *gin.Context) {
			prod := &Product{}
			c.BindJSON(prod)
			Update(db, c.Param("id"), prod.Name, prod.Price)
			c.JSON(http.StatusOK, gin.H{
				"message": "Product updated",
			})
		})
	}

	r.Run(":3000")
}

func GetAll(db *sql.DB) []Product {
	var prod Product
	productTab := []Product{}

	rows, err := db.Query(`SELECT * FROM product`)
	CheckError(err)

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&prod.Id, &prod.Name, &prod.Price, &prod.CreatedAt, &prod.UpdatedAt)
		CheckError(err)
		productTab = append(productTab, prod)
	}

	return productTab
}

func GetById(db *sql.DB, id string) Product {
	var prod Product

	rows, err := db.Query(`SELECT * FROM product where id=$1`, id)
	CheckError(err)

	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&prod.Id, &prod.Name, &prod.Price, &prod.CreatedAt, &prod.UpdatedAt)
		CheckError(err)
	}

	return prod
}

func Delete(db *sql.DB, id string) {
	_, err := db.Exec(`DELETE FROM product where id=$1`, id)
	CheckError(err)
}

func Update(db *sql.DB, id string, name string, price float32) {
	_, err := db.Exec(`UPDATE product SET name=$1, price=$2, updatedAt=now() WHERE id=$3`, name, price, id)
	CheckError(err)
}

func Create(db *sql.DB, name string, price float32) {
	_, err := db.Exec(`INSERT INTO product (name, price, createdAt, updatedAt) VALUES ($1, $2, now(), now())`, name, price)
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
