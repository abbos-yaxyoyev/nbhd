package database

import "nbhd/models"

type taskStorage interface {
	CreateTask(models.Task) error
}
