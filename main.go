package main

import (
	"fmt"
	"log"
	"net/http"
	"handlers"
	"github.com/gorilla/mux"
)


func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user", handlers.CreateUserEP).Methods("POST")
	router.HandleFunc("/user", handlers.GetAllUserEP).Methods("GET")
	router.HandleFunc("/user/{email}", handlers.GetUserEP).Methods("GET")
	fmt.Println("Starting API :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}	