package main

import (
	"flag"

	"nbhd/env/config"
	"nbhd/env/database/postgres"
	"nbhd/env/server"
	"nbhd/env/transfer/api"
	"nbhd/env/validator/validator"
	"nbhd/tools/logger"
	"nbhd/usecases"
)

func main() {

	envName := *flag.String("c", "default.cfg", "Environment config name")

	flag.Parse()

	cfg, err := config.NewConfig(envName)

	if err != nil {
		logger.Error(err.Error())
	}

	db, err := postgres.NewPostgresDatabase(cfg.Db)

	if err != nil {
		logger.Error(err.Error())
	}

	validate := validator.NewValidator()

	controller := usecases.NewController(db, validate)

	APIHandler := api.NewAPIHandler(cfg.Handler, controller)

	err = server.RunServer(cfg.Server, APIHandler)

	if err != nil {
		logger.Error(err.Error())
	}

}
