package database

import (
	"os"

	"github.com/jinzhu/gorm"
)

var Conn *gorm.DB

func InitDB() (Conn *gorm.DB) {
	dbuser := os.Getenv("DBUSER")
	dbpassword := os.Getenv("DBPASSWORD")
	dbname := os.Getenv("DBNAME")
	Conn, err := gorm.Open("mysql", dbuser+":"+dbpassword+"@/"+dbname+"?charset=utf8&parseTime=True")
	if err != nil {
		panic("failed to connect database")
	}
	Conn.DB().Ping()
	return
}
