package handlers

import (
	"encoding/json"
	"net/http"
	"models"
	"function"
	"fmt"
	"github.com/gorilla/mux"
)

func CreateUserEP(w http.ResponseWriter, r *http.Request) {
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	function.InsertOne(user)
	json.NewEncoder(w).Encode(user)
}

func GetAllUserEP(w http.ResponseWriter, r *http.Request) {
	data := function.GetAllUser()
	json.NewEncoder(w).Encode(data)
}

func GetUserEP(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	data := function.GetAllUser()
	for _, p := range data {
		if p.PersonalInfo.Email == params["email"] {
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	json.NewEncoder(w).Encode("User not found")
}

func DeleteUserEP(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	data := function.GetAllUser()
	for _, p := range data {
		if p.PersonalInfo.Email == params["email"] {
			function.DeleteUser(p)
			fmt.Fprintf(w, "Delete User Success")
			return 
		}
	}
	json.NewEncoder(w).Encode("User not found")
}