package routes

import (
	"net/http"
	"workmate/db"
	"workmate/handlers"

	"github.com/gorilla/mux"
)

func InitRoutes(r *mux.Router, manager *db.Manager) {
	CreateTask(r, manager)
	CreateUser(r, manager)
	GetUser(r, manager)
	GetTask(r, manager)
	GetTasks(r, manager)
}

func CreateUser(r *mux.Router, manager *db.Manager) {
	r.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateUser(w, r, manager)
	}).Methods("POST")
}

func GetUser(r *mux.Router, manager *db.Manager) {
	r.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetUser(w, r, manager)
	}).Methods("GET")
}

func GetTask(r *mux.Router, manager *db.Manager) {
	r.HandleFunc("/task/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTask(w, r, manager)
	}).Methods("GET")
}

func GetTasks(r *mux.Router, manager *db.Manager) {
	r.HandleFunc("/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllTasks(w, r, manager)
	}).Methods("GET")
}

func CreateTask(r *mux.Router, manager *db.Manager) {
	r.HandleFunc("/task/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateTask(w, r, manager)
	}).Methods("POST")
}
