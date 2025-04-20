package main

import (
	"fmt"
	"workmate/db"
	"workmate/routes"
	"workmate/services"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	/*
		Initialization
	*/
	services.LoadEnvFile(".env")
	manager := db.NewDBManager("postgres", os.Getenv("DB_CONNECTION_STRING"))
	db.Migrate(manager)

	log.Println("База данных успешно инициализирована и мигрирована.")
	defer manager.Close()

	r := mux.NewRouter()

	routes.InitRoutes(r, manager)

	/*
		Server initialization
	*/
	serverPort := os.Getenv("SERVER_PORT")
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", serverPort),
		Handler: r,
	}

	log.Println("Сервер запущен на порту " + serverPort)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}