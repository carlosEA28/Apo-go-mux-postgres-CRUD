package main

import (
	"net/http"

	"github.com/carlos/go-mux-postgres/db"
	"github.com/carlos/go-mux-postgres/models"
	"github.com/carlos/go-mux-postgres/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConnection()                //inicia o banco de dados
	db.DB.AutoMigrate(models.Task{}) //gera as migrações para a criação das tabelas tarefas
	db.DB.AutoMigrate(models.User{}) //gera as migrações para a criação das tabelas users

	r := mux.NewRouter() //inicia o router

	r.HandleFunc("/", routes.HomeHandler)

	//Rotas User
	r.HandleFunc("/user/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users", routes.CreateUsersHandler).Methods("POST")
	r.HandleFunc("/user/{id}", routes.DeleteUsersHandler).Methods("DELETE")

	//Rotas Tasks
	r.HandleFunc("/task/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks", routes.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/task/{id}", routes.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}
