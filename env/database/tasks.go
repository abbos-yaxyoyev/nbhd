package database

import "nbhd/models"

type taskStorage interface {
	CreateTask(models.Task) error
	GetTask(string) (models.Task, error)
	ListTasks([4]float64) ([]models.Task, error)
}
