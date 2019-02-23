package handlers

import (
	"encoding/json"
	"net/http"
	"models"
	"function"
	// "github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	function.InsertOne(user)
	json.NewEncoder(w).Encode(user)
}
