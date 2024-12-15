package endpoint

import (
	"assesment-sigmatech/config/logging"
	"assesment-sigmatech/service/api"
	"assesment-sigmatech/service/middleware"
	"assesment-sigmatech/service/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Endpoint(apiRoute *api.AccountApi, ds *repository.DatabaseData, logger *logging.Logger) *echo.Echo {
	e := echo.New()

	e.POST("/register", apiRoute.Register)

	middleware := middleware.InitMiddleWare(*ds, logger)

	protected := e.Group("/v2")
	protected.Use(middleware.BasicAuthMiddleWare)
	protected.POST("/create-account", apiRoute.CreateAccount)
	protected.POST("/upload-id-photo", apiRoute.UploadIDPhoto)
	protected.POST("/upload-selfie-photo", apiRoute.UploadIDPhoto)
	// protected.POST("/tarik", apiRoute.Tarik)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	return e
}
