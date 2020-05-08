package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	db "github.com/dariuszkorolczukcom/pinkbubbleapi/database"
	m "github.com/dariuszkorolczukcom/pinkbubbleapi/models"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {

	var categories []m.Category
	db.Conn.Find(&categories)
	fmt.Println(categories)

	json.NewEncoder(w).Encode(categories)
}
