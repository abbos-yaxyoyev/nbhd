package database

type Database interface {
	usersStorage
	sessionsStorage
	taskStorage
}
