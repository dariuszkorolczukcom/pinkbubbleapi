package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email     string `json:"email"`
	Role      int    `json:"status"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Category struct {
	gorm.Model
	Name     string    `json:"name"`
	Products []Product `gorm:"foreignkey:CategoryId" json:"products"`
}

func (u Category) TableName() string {
	return "categories"
}

type Product struct {
	gorm.Model
	Img              string  `json:"img"`
	Price            float64 `json:"price"`
	Name             string  `json:"name"`
	ShortDescription string  `json:"shortDescription"`
	Description      string  `json:"description"`
	CategoryId       uint    `json:"categoryID"`
	Images           []Image `gorm:"foreignkey:ProductId" json:"images"`
}

type Image struct {
	gorm.Model
	Name      string `json:"name"`
	ProductId uint   `json:"productID"`
}
