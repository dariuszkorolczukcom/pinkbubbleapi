package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

var Conn *gorm.DB

func InitDB() (Conn *gorm.DB) {
	fmt.Println("initialising DB")
	dbuser := os.Getenv("DBUSER")
	dbpassword := os.Getenv("DBPASSWORD")
	dbname := os.Getenv("DBNAME")
	dbhost := os.Getenv("DBHOST")
	dbport := os.Getenv("DBPORT")
	fmt.Println("mysql", dbuser+":"+dbpassword+"@("+dbhost+":"+dbport+")/"+dbname+"?charset=utf8&parseTime=true")
	Conn, err := gorm.Open("mysql", dbuser+":"+dbpassword+"@("+dbhost+":"+dbport+")/"+dbname+"?charset=utf8&parseTime=true")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("DB initiallised")
	Conn.LogMode(true)
	return
}
