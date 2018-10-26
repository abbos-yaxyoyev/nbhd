package postgres

import (
	"github.com/lib/pq"
	"nbhd/models"
	"nbhd/tools/logger"
)

func (db Database) CreateTask(task models.Task) error {

	query := "INSERT INTO tasks(id, title, category, location, description, time, creator, performer, encouragement, pay, created, archived) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)"

	_, err := db.db.Exec(query, task.Id, task.Title, task.Category, pq.Array(task.Location), task.Description, task.Time, task.Creator, task.Performer, task.Encouragement, task.Pay, task.Created, task.Archived)

	if err != nil {
		logger.Warning(err.Error())
		return err
	}

	return nil

}
