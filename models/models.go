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

type Order struct {
	gorm.Model
	UserId     uint
	AddressId  uint
	Status     int         `json:"status"`
	Price      float64     `json:"price"`
	OrderItems []OrderItem `gorm:"foreignkey:OrderId" json:"orderItems"`
	User       User        `gorm:"foreignkey:UserId" json:"user"`
	Address    Address     `gorm:"foreignkey:AddressId" json:"address"`
}

type OrderItem struct {
	gorm.Model
	OrderId   uint
	ProductId uint
	Quantity  int     `json:"qty"`
	Price     float64 `json:"pricePerItem"`
	Product   Product `gorm:"foreignkey:ProductId" json:"product"`
}

type Address struct {
	gorm.Model
	UserId   uint
	Name     string `json:"name"`
	Email    string `json:"email"`
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	Postcode string `json:"postcode"`
	Town     string `json:"town"`
}

func (u Address) TableName() string {
	return "address"
}

type Faq struct {
	gorm.Model
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func (u Faq) TableName() string {
	return "faq"
}
