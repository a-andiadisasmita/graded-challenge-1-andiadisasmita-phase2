package main

import (
	"graded-challenge-1-andiadisasmita/database"
	"graded-challenge-1-andiadisasmita/handlers"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// Initialize the database connection
	db := database.Connect()
	defer db.Close()

	// Create a new router
	router := httprouter.New()

	// Define routes
	router.GET("/customers", handlers.GetAllCustomers(db))
	router.GET("/customers/:id", handlers.GetCustomerByID(db))
	router.POST("/customers", handlers.CreateCustomer(db))
	router.PUT("/customers/:id", handlers.UpdateCustomer(db))
	router.DELETE("/customers/:id", handlers.DeleteCustomer(db))

	// Start the server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
