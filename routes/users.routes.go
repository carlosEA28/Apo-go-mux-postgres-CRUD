package routes

import (
	"encoding/json"
	"net/http"

	"github.com/carlos/go-mux-postgres/db"
	"github.com/carlos/go-mux-postgres/models"
	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"]) // vai encontrar o primeiro registro que tenha o id passado pelo parametro(nos usamos os colchetes pois o nux.vars retorna um map)

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return // para acabar a funcao, um return vazio
	}

	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)

	json.NewEncoder(w).Encode(&user)

}

func CreateUsersHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	createUser := db.DB.Create(&user)
	err := createUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&user)
}

func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return // para acabar a funcao, um return vazio
	}

	db.DB.Unscoped().Delete(&user)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted"))
}
