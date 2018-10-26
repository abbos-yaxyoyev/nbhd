package database

import "nbhd/models"

type usersStorage interface {
	GetUser(string) (models.User, error)
}
