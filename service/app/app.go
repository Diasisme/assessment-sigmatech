package app

import "assesment-sigmatech/config/logging"

type accountApp struct {
	accRepo AccountDatastore
	log     *logging.Logger
}

func InitApp(db AccountDatastore, log *logging.Logger) AccountApp {
	return &accountApp{
		accRepo: db,
		log:     log,
	}
}
