package middleware

import (
	"assesment-sigmatech/config/logging"
	"assesment-sigmatech/config/logging/utils"
	"assesment-sigmatech/service/models"
	"assesment-sigmatech/service/repository"
	"net/http"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type MiddlewareApp struct {
	database *repository.DatabaseData
	log      *logging.Logger
}

func (m *MiddlewareApp) BasicAuthMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		account_number := c.Request().Header.Get("X-Account_number")
		pin := c.Request().Header.Get("X-Pin")
		db := m.database.DB
		if account_number == "" || pin == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Username and password required"})
		}

		m.log.Info(logrus.Fields{
			"account_number": account_number,
		}, nil, "")

		var data models.UserLogin
		if err := db.Where("account_number", account_number).First(&data).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				m.log.Error(logrus.Fields{"err": err}, nil, "error: Invalid username or password")
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"})
			}
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Database error"})
		}

		if !utils.CheckPinHash(data.Pin, pin) {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"})
		}

		c.Set("account_number:", data.AccountNumber)
		return next(c)
	}
}

func InitMiddleWare(db repository.DatabaseData, logger *logging.Logger) *MiddlewareApp {
	return &MiddlewareApp{
		database: &db,
		log:      logger,
	}
}
