package postgres

import (
	"database/sql"
	"github.com/lib/pq"
	"nbhd/models"
	"nbhd/tools/logger"
)

func (db Database) GetUser(id string) (models.User, error) {

	var user models.User

	query := "SELECT id, name, phone, photo, location, password FROM users WHERE id = $1"

	err := db.db.QueryRow(query, id).Scan(&user.Id, &user.Name, &user.Phone, &user.Photo, pq.Array(&user.Location), &user.Password)

	if err != nil && err != sql.ErrNoRows {
		logger.Warning(err.Error())
		return user, err
	}

	return user, nil
}

func (db Database) GetUserByPhone(id string) (models.User, error) {

	var user models.User

	query := "SELECT id, name, phone, photo, location, password FROM users WHERE phone = $1"

	err := db.db.QueryRow(query, id).Scan(&user.Id, &user.Name, &user.Phone, &user.Photo, pq.Array(&user.Location), &user.Password)

	if err != nil && err != sql.ErrNoRows {
		logger.Warning(err.Error())
		return user, err
	}

	return user, nil
}

func (db Database) StoreUser(user models.User) error {

	if user.Location == nil {
		user.Location = make([]float64, 0)
	}

	query := "INSERT INTO users(id, name, photo, phone, location, password) VALUES ($1, $2, $3, $4, $5, $6)"

	_, err := db.db.Exec(query, user.Id, user.Name, user.Photo, user.Phone, pq.Array(&user.Location), user.Password)

	if err != nil {
		logger.Warning(err.Error())
		return err
	}

	return nil

}

func (db Database) UpdateUser(user models.User) error {

	if user.Location == nil {
		user.Location = make([]float64, 0)
	}

	query := "UPDATE users SET name = $2, photo = $3, phone = $4, location = $5, password = $6 WHERE id = $1"

	_, err := db.db.Exec(query, user.Id, user.Name, user.Photo, user.Phone, pq.Array(&user.Location), user.Password)

	if err != nil {
		logger.Warning(err.Error())
		return err
	}

	return nil

}

func (db Database) GetUserRating(id string) (float64, error) {

	var count1 float64
	var count2 float64

	query := "SELECT SUM(creator_rating)/COUNT(creator_rating) FROM tasks WHERE creator = $1 AND creator_rating != 0"

	err := db.db.QueryRow(query, id).Scan(&count1)

	if err != nil {
		logger.Warning(err.Error())
		return 0, err
	}

	query = "SELECT SUM(performer_rating)/COUNT(performer_rating) FROM tasks WHERE performer = $1 AND performer_rating != 0"

	err = db.db.QueryRow(query, id).Scan(&count2)

	if err != nil {
		logger.Warning(err.Error())
		return 0, err
	}

	return (count1 + count2) / 2, nil

}
