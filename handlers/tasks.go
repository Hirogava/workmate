package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"workmate/db"
	"workmate/tasks"

	"github.com/gorilla/mux"
)

func GetTask(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Println("Неверный формат ответа")
		http.Error(w, "Неверный формат ответа", http.StatusBadRequest)
		return
	}

	task, err := manager.GetTask(id)
	if err != nil {
		log.Printf("Ошибка получения задачи: %s", err)
		http.Error(w, "Ошибка получения задачи", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func GetAllTasks(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("Неверный формат ответа")
		http.Error(w, "Неверный формат ответа", http.StatusBadRequest)
		return
	}

	tasks, err := manager.GetTasks(userId)
	if err != nil {
		log.Printf("Ошибка получения задач: %s", err)
		http.Error(w, "Ошибка получения задач", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("Неверный формат ответа")
		http.Error(w, "Неверный формат ответа", http.StatusBadRequest)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		log.Println("Неверный формат ответа")
		http.Error(w, "Content-Type должен быть application/json", http.StatusBadRequest)
		return
	}

	var request struct {
		Task string `json:"task"`
	}

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Printf("Неверный формат JSON: %v", err)
		http.Error(w, "Неверный формат JSON", http.StatusBadRequest)
		return
	}

	task, err := manager.CreateTask(userId, "pending", request.Task)
	if err != nil {
		log.Printf("Ошибка создания задачи: %s", err)
		http.Error(w, "Ошибка создания задачи", http.StatusInternalServerError)
		return
	}

	tasks.RenderTask(task, manager)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Таск с id " + strconv.Itoa(task.ID) + " создан"})
}