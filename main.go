package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	InitDB("user=postgres password=1009 dbname=auth_service sslmode=disable")

	router := mux.NewRouter()

	router.HandleFunc("/token", TokenHandler).Methods("GET")
	router.HandleFunc("/refresh", RefreshHandler).Methods("POST")

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
