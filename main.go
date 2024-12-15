package main

import (
	"assesment-sigmatech/config"
	"assesment-sigmatech/config/logging"
	"assesment-sigmatech/service/api"
	"assesment-sigmatech/service/app"
	"assesment-sigmatech/service/middleware"
	"assesment-sigmatech/service/repository"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	viper := config.NewViper()

	varenv := config.NewEnvVar(viper)
	logger := logging.NewLogger(varenv.Service)

	ds := repository.InitDB(varenv, logger)
	appRoute := app.InitApp(ds, logger)
	apiRoute := api.InitApi(appRoute, logger)

	e.POST("/register", apiRoute.Register)

	middleware := middleware.InitMiddleWare(*ds, logger)

	protected := e.Group("/v2")
	protected.Use(middleware.BasicAuthMiddleWare)
	protected.POST("/create-account", apiRoute.CreateAccount)
	// protected.POST("/tarik", apiRoute.Tarik)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8100"))

	e.Start(fmt.Sprintf(":%s", varenv.ServicePort))
}
