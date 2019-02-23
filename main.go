package main

import (
	"fmt"
	"log"
	"net/http"
	"handlers"
	"github.com/gorilla/mux"
	// "github.com/mongodb/mongo-go-driver/mongo"
)


func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user", handlers.CreateUser).Methods("POST")
	fmt.Println("Starting API :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}