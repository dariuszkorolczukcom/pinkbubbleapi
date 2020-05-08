package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name string `json:"name"`
}

func (u Category) TableName() string {
	return "categories"
}

type Product struct {
	gorm.Model
	Img               string  `json:"img"`
	Price             float64 `json:"price"`
	Name              string  `json:"title"`
	Short_description string  `json:"subtitle"`
	Description       string  `json:"text"`
	CategoryId        int
	Category          Category `gorm:"foreignkey:CategoryId"`
}

// func (u Product) TableName() string {
// 	return "products"
// }
