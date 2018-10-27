package postgres

import (
	"database/sql"
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

func (db Database) UpdateTask(task models.Task) error {

	query := "UPDATE tasks SET title = $2, category = $3, location = $4, description = $5, time = $6, creator = $7, performer = $8, encouragement = $9, pay = $10, created = $11, archived = $12 WHERE id = $1"

	_, err := db.db.Exec(query, task.Id, task.Title, task.Category, pq.Array(task.Location), task.Description, task.Time, task.Creator, task.Performer, task.Encouragement, task.Pay, task.Created, task.Archived)

	if err != nil {
		logger.Warning(err.Error())
		return err
	}

	return nil

}

func (db Database) GetTask(id string) (models.Task, error) {

	var task models.Task

	query := "SELECT id, title, category, location, description, time, creator, performer, encouragement, pay, created, archived FROM tasks WHERE id = $1"

	err := db.db.QueryRow(query, id).Scan(&task.Id, &task.Title, &task.Category, pq.Array(&task.Location), &task.Description, &task.Time, &task.Creator, &task.Performer, &task.Encouragement, &task.Pay, &task.Created, &task.Archived)

	if err != nil && err != sql.ErrNoRows {
		logger.Warning(err.Error())
		return task, err
	}

	return task, nil
}

func (db Database) ListTasks([4]float64) ([]models.Task, error) {

	tasks := make([]models.Task, 0)

	query := "SELECT id, title, category, location, description, time, creator, performer, encouragement, pay, created, archived FROM tasks WHERE archived = FALSE ORDER BY created DESC"

	rows, err := db.db.Query(query)

	if err != nil && err != sql.ErrNoRows {
		logger.Warning(err.Error())
		return tasks, err
	}

	defer rows.Close()

	for rows.Next() {
		var task models.Task
		rows.Scan(&task.Id, &task.Title, &task.Category, pq.Array(&task.Location), &task.Description, &task.Time, &task.Creator, &task.Performer, &task.Encouragement, &task.Pay, &task.Created, &task.Archived)
		tasks = append(tasks, task)
	}

	return tasks, nil

}

func (db Database) GetTaskPerformer(task, user string) (models.TaskPerformer, error) {

	var performer models.TaskPerformer

	query := "SELECT task_id, user_id FROM task_performers WHERE task_id = $1 AND user_id = $2"

	err := db.db.QueryRow(query, task, user).Scan(&performer.Task, &performer.User)

	if err != nil && err != sql.ErrNoRows {
		logger.Warning(err.Error())
		return performer, err
	}

	return performer, nil
}

func (db Database) StoreTaskPerformer(performer models.TaskPerformer) error {

	query := "INSERT INTO task_performers (task_id, user_id) VALUES ($1, $2)"

	_, err := db.db.Exec(query, performer.Task, performer.User)

	if err != nil {
		logger.Warning(err.Error())
		return err
	}

	return nil

}

func (db Database) DeleteTaskPerformer(performer models.TaskPerformer) error {

	query := "DELETE FROM task_performers WHERE task_id = $1 AND user_id = $2"

	_, err := db.db.Exec(query, performer.Task, performer.User)

	if err != nil {
		logger.Warning(err.Error())
		return err
	}

	return nil

}
