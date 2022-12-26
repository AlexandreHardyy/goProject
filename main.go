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

type product struct {
	id        int
	name      string
	price     int
	createdAt string
	updatedAt string
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

	r.Run(":3000")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
