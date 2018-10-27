package database

import "nbhd/models"

type taskStorage interface {
	CreateTask(models.Task) error
	UpdateTask(models.Task) error
	GetTask(string) (models.Task, error)
	ListTasks([4]float64) ([]models.Task, error)
	GetTaskPerformer(string, string) (models.TaskPerformer, error)
	StoreTaskPerformer(models.TaskPerformer) error
	DeleteTaskPerformer(models.TaskPerformer) error
	ListTaskPerformers(string) ([]models.TaskPerformer, error)
}
