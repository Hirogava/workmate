package db

import "workmate/models"

func (Manager *Manager) CreateUser(name string) (models.User, error) {
	var user models.User
	query := `INSERT INTO users (name) VALUES ($1) RETURNING id`
	err := Manager.Conn.QueryRow(query, name).Scan(&user.ID)
	if err != nil {
		return models.User{}, err
	}
	user.Name = name
	return user, nil
}

func (Manager *Manager) GetUser(id int) (models.User, error) {
	var user models.User
	query := `SELECT id, name FROM users WHERE id = $1`
	err := Manager.Conn.QueryRow(query, id).Scan(&user.ID, &user.Name)
	if err != nil {
		return models.User{}, err
	}
    return user, nil
}