package main

import (
	"assesment-sigmatech/config"
	"assesment-sigmatech/config/logging"
	"assesment-sigmatech/endpoint"
	"assesment-sigmatech/service/api"
	"assesment-sigmatech/service/app"
	"assesment-sigmatech/service/minio"
	"assesment-sigmatech/service/repository"
)

func main() {

	viper := config.NewViper()

	varenv := config.NewEnvVar(viper)
	logger := logging.NewLogger(varenv.Service)

	minioService := minio.InitMinio(varenv, logger)
	ds := repository.InitDB(varenv, logger)
	appRoute := app.InitApp(ds, logger, minioService)
	apiRoute := api.InitApi(appRoute, logger)

	e := endpoint.Endpoint(apiRoute, ds, logger)
	e.Logger.Fatal(e.Start(":8100"))
}
