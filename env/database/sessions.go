package database

import "nbhd/models"

type sessionsStorage interface {
	GetSession(string) (models.Session, error)
}
