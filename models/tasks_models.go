package models

import "time"

type Task struct {
	ID        int  `json:"id"`
	Status    string `json:"status"`
	Task      string `json:"task"`
	Result    string `json:"result"`
	Error     string `json:"error"`
	CreatedAt string `json:"createdAt"`
	EndedAt   time.Time `json:"endedAt"`
	User_id   int64  `json:"userId"`
}