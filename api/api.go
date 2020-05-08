package main

import (
	"fmt"
	"log"
	"net/http"

	db "github.com/dariuszkorolczukcom/pinkbubbleapi/database"
	h "github.com/dariuszkorolczukcom/pinkbubbleapi/handlers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I am an API!\n")
	fmt.Fprintf(w, "Build by an intelligent being,\n")
	fmt.Fprintf(w, "Running on a stupid machine.")
}

func main() {
	db.Conn = db.InitDB()

	defer db.Conn.Close()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	// products
	router.HandleFunc("/products", h.GetProducts).Methods("GET")
	router.HandleFunc("/product", h.GetProduct).Methods("GET")
	router.HandleFunc("/product", h.AddProduct).Methods("POST")
	router.HandleFunc("/product", h.UpdateProduct).Methods("PUT")
	router.HandleFunc("/product", h.DeleteProduct).Methods("DELETE")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	//categories
	router.HandleFunc("/categories", h.GetCategories)
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
