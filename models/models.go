package models

import (
	"time"
)

type Product struct {
	Id        int       `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" binding:"required"`
	Price     float32   `json:"price" binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Payment struct {
	Id        int       `json:"id" gorm:"primary_key"`
	ProductID int       `json:"product_id"`
	Product   Product   `gorm:"foreignKey:ProductID" gorm:"references:ProductID"`
	PricePaid float64   `json:"pricePaid"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
