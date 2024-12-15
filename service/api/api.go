package api

import (
	"assesment-sigmatech/config/logging"
	"assesment-sigmatech/service/app"

	v1 "github.com/go-playground/validator/v10"
)

type AccountApi struct {
	app      app.AccountApp
	validate v1.Validate
	log      *logging.Logger
}

func InitApi(app app.AccountApp, log *logging.Logger) *AccountApi {
	return &AccountApi{
		app:      app,
		validate: *v1.New(),
		log:      log,
	}
}
