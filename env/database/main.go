package database

import "nbhd/models"

type Database interface {
	GetUserById(int) (models.User, error)
}
