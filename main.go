package main

import (
	"AuthManagement/config"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

	func main() {
		
		// database connection

		dbConfig := &config.DatabaseConfig{
			URI:      "mongodb://localhost:27017",
			Databse: "go-auth-db",
			Collections: []string{
				"users",
				"otp",
			},
	
		}
		databse, err := config.Connect(dbConfig)

		if err != nil {
			log.Fatal(err)
		}
		
		fmt.Println("database is connected",databse)

		router := mux.NewRouter()
		port := ":8080"

		fmt.Println("server connected with port", port)

		log.Fatal(http.ListenAndServe(port, router))

	}