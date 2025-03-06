package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go-crud/database"
	"go-crud/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	database.DB.Create(&user)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	database.DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User
	if database.DB.First(&user, params["id"]).Error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Usuário não encontrado")
		return
	}
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User
	if database.DB.First(&user, params["id"]).Error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Usuário não encontrado")
		return
	}
	json.NewDecoder(r.Body).Decode(&user)
	database.DB.Save(&user)
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User
	if database.DB.First(&user, params["id"]).Error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Usuário não encontrado")
		return
	}
	database.DB.Delete(&user)
	w.WriteHeader(http.StatusNoContent)
}
