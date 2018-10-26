package postgres

import (
	"database/sql"
	"nbhd/models"
	"nbhd/tools/logger"
)

func (db Database) GetSession(id string) (models.Session, error) {

	var session models.Session

	query := "SELECT id, user_id FROM sessions WHERE id = $1"

	err := db.db.QueryRow(query, id).Scan(&session.Id, &session.User)

	if err != nil && err != sql.ErrNoRows {
		logger.Warning(err.Error())
		return session, err
	}

	return session, nil
}
