package usecases

import (
	"nbhd/env/database"
	"nbhd/env/validator"
)

type Controller struct {
	db        database.Database
	validator validator.Validator
}

func NewController(db database.Database, validator validator.Validator) Controller {
	return Controller{db, validator}
}
