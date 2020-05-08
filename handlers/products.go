package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	db "github.com/dariuszkorolczukcom/pinkbubbleapi/database"
	m "github.com/dariuszkorolczukcom/pinkbubbleapi/models"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	var product m.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.Conn.First(&product)
	fmt.Println(product)

	json.NewEncoder(w).Encode(product)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {

	var products []m.Product
	db.Conn.Find(&products)

	json.NewEncoder(w).Encode(products)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {

	var product m.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.Conn.Create(&product)
	db.Conn.Find(&product)

	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {

	var product m.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.Conn.Save(&product)
	db.Conn.First(&product)

	json.NewEncoder(w).Encode(product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var product m.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db.Conn.Delete(&product)
	db.Conn.First(&product)
	fmt.Println(product)

	json.NewEncoder(w).Encode(product)
}

// [
//     {
//         "name": "bath",
//         "items": [
//             {
//                 "id": 1,
//                 "img": "bomb.jpg",
//                 "price": 4.99,
//                 "title": "card title",
//                 "subtitle": "card subtitle",
//                 "text":
//                     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam at massa ac leo congue tincidunt vel at velit."
//             },
//             {
//                 "id": 2,
//                 "category": 1,
//                 "img": "bomb.jpg",
//                 "price": 4.99,
//                 "title": "card title 2",
//                 "subtitle": "card subtitle",
//                 "text":
//                     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam at massa ac leo congue tincidunt vel at velit."
//             },
//             {
//                 "id": 3,
//                 "category": 1,
//                 "img": "bomb.jpg",
//                 "price": 4.99,
//                 "title": "card title 3",
//                 "subtitle": "card subtitle",
//                 "text":
//                     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam at massa ac leo congue tincidunt vel at velit."
//             },
//             {
//                 "id": 3,
//                 "category": 1,
//                 "title": "card title 3",
//                 "price": 4.99,
//                 "subtitle": "card subtitle",
//                 "text":
//                     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam at massa ac leo congue tincidunt vel at velit."
//             }
//         ]
//     },
//     {
//         "name": "beauty",
//         "items": [
//             {
//                 "id": 3,
//                 "category": 2,
//                 "title": "card title 3",
//                 "price": 4.99,
//                 "subtitle": "card subtitle",
//                 "text":
//                     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam at massa ac leo congue tincidunt vel at velit."
//             },
//             {
//                 "id": 3,
//                 "category": 2,
//                 "title": "card title 3",
//                 "subtitle": "card subtitle",
//                 "text":
//                     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam at massa ac leo congue tincidunt vel at velit."
//             },
//             {
//                 "id": 3,
//                 "category": 2,
//                 "title": "card title 3",
//                 "subtitle": "card subtitle",
//                 "text":
//                     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam at massa ac leo congue tincidunt vel at velit."
//             },
//             {
//                 "id": 3,
//                 "category": 2,
//                 "title": "card title 3",
//                 "subtitle": "card subtitle",
//                 "text":
//                     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam at massa ac leo congue tincidunt vel at velit."
//             },
//             {
//                 "id": 3,
//                 "category": 2,
//                 "title": "card title 3",
//                 "subtitle": "card subtitle",
//                 "text":
//                     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam at massa ac leo congue tincidunt vel at velit."
//             },
//             {
//                 "id": 3,
//                 "category": 2,
//                 "title": "card title 3",
//                 "subtitle": "card subtitle",
//                 "text":
//                     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam at massa ac leo congue tincidunt vel at velit."
//             },
//             {
//                 "id": 3,
//                 "category": 2,
//                 "title": "card title 3",
//                 "subtitle": "card subtitle",
//                 "text":
//                     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam at massa ac leo congue tincidunt vel at velit."
//             }
//         ]
//     },
//     {
//         "name": "perfumes",
//         "items": [
//             {
//                 "id": 3,
//                 "category": 3,
//                 "title": "card title 3",
//                 "subtitle": "card subtitle",
//                 "text":
//                     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam at massa ac leo congue tincidunt vel at velit."
//             }
//         ]
//     }
// ]
