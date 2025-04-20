package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"workmate/db"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	if r.Header.Get("Content-Type") != "application/json" {
		log.Println("Неверный формат ответа")
		http.Error(w, "Content-Type должен быть application/json", http.StatusBadRequest)
		return
	}

	var request struct {
		Name string `json:"name"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Printf("Неверный формат JSON: %v", err)
		http.Error(w, "Неверный формат JSON", http.StatusBadRequest)
		return
	}

	user, err := manager.CreateUser(request.Name)
	if err != nil {
		log.Printf("Ошибка создания пользователя: %v", err)
		http.Error(w, "Ошибка создания пользователя", http.StatusInternalServerError)
		return
	}

	var response struct {
		ID int64 `json:"id"`
		Name string `json:"name"`
	}

	response.ID, response.Name = user.ID, user.Name

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	log.Printf("Пользователь %s создан с ID %d", user.Name, user.ID)
}

func GetUser(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("Неверный формат ответа")
		http.Error(w, "Неверный формат ответа", http.StatusBadRequest)
		return
	}

	user, err := manager.GetUser(userID)
	if err != nil {
		log.Printf("Ошибка получения пользователя: %v", err)
		http.Error(w, "Ошибка получения пользователя", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}