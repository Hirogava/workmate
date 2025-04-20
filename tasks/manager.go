package tasks

import (
	"crypto/rand"
	"log"
	"math/big"
	"time"
	"workmate/db"
	"workmate/models"
)

func RenderTask(task models.Task, manager *db.Manager) {
	go func() {
		log.Printf("Render task: %s", task.Task)
		// имитация долгой работы таска
		time.Sleep(time.Second * 3)
		statuses := []string{"done", "failed"}
		// имитация результата и ошибки
		results := []string{"1234567890", "error"}

		intStat, err := rand.Int(rand.Reader, big.NewInt(int64(len(statuses))))
		if err != nil {
			task.Error = err.Error()
			task.Status = "failed"
			return
		}

		task.Status = statuses[intStat.Int64()]
		if task.Status == "failed" {
			task.Result = results[1]
		} else {
			task.Result = results[0]
		}

		task.EndedAt = time.Now()
		err = manager.EndTask(task)
		if err != nil {
			log.Printf("Ошибка выполнения таска: %v", err)
		}

		log.Printf("Таск с ID %d завершен", task.ID)
	}()
}
