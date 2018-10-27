package database

import "nbhd/models"

type usersStorage interface {
	GetUser(string) (models.User, error)
	GetUserByPhone(string) (models.User, error)
	StoreUser(models.User) error
	UpdateUser(models.User) error
}
