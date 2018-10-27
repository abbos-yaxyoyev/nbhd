package database

import "nbhd/models"

type sessionsStorage interface {
	StoreSession(models.Session) error
	GetSession(string) (models.Session, error)
	DeleteSession(models.Session) error
}
