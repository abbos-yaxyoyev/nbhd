package postgres

import (
	"database/sql"
	"github.com/lib/pq"
	"nbhd/models"
	"nbhd/tools/logger"
)

func (db Database) GetUser(id string) (models.User, error) {

	var user models.User

	query := "SELECT id, name, phone, photo, location FROM users WHERE id = $1"

	err := db.db.QueryRow(query, id).Scan(&user.Id, &user.Name, &user.Phone, &user.Photo, pq.Array(&user.Location))

	if err != nil && err != sql.ErrNoRows {
		logger.Warning(err.Error())
		return user, err
	}

	return user, nil
}
