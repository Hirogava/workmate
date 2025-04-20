package db

import (
	"workmate/models"
)

func (manager *Manager) GetTask(id int) (models.Task, error) {
	var task models.Task
	query := "SELECT * FROM tasks WHERE id=$1"
	err := manager.Conn.QueryRow(query, id).Scan(&task.ID, &task.Status, &task.Task, &task.Result,
		&task.Error, &task.CreatedAt, &task.EndedAt, task.User_id)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (manager *Manager) GetTasks(userId int) ([]models.Task, error) {
	var tasks []models.Task

	query := "SELECT * FROM tasks WHERE user_id=$1"
	rows, err := manager.Conn.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
			var task models.Task
			err := rows.Scan(&task.ID, &task.Status, &task.Task,
				&task.Result, &task.Error, &task.CreatedAt, &task.User_id)
			if err != nil {
				return nil, err
			}
			tasks = append(tasks, task)
	}

   return tasks, nil
}

func (manager *Manager) CreateTask(userId int, status string, quest string) (models.Task, error) {
	var task models.Task

	query := "INSERT INTO tasks (user_id, status, task) VALUES ($1, $2, $3) RETURNING id"
	err := manager.Conn.QueryRow(query, userId, status, quest).Scan(&task.ID)
	if err != nil {
		return models.Task{}, err
	}

	task.Status, task.Task = status, quest

    return task, nil
}

func (manager *Manager) EndTask(task models.Task) error {
	query := "UPDATE tasks SET status=$1, result=$2, error=$3, ended_at=$4 WHERE id=$5"
	_, err := manager.Conn.Exec(query, task.Status, task.Result, task.Error, task.EndedAt, task.ID)
	return err
}