package repository

import (
	"assesment-sigmatech/config/logging"
	"assesment-sigmatech/service/models"
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseData struct {
	DB  *gorm.DB
	log *logging.Logger
}

func (f *DatabaseData) Begin() (*gorm.DB, error) {
	var err error

	tx := f.DB.Begin()
	if err != nil {
		f.log.Error(logrus.Fields{
			"error": err.Error(),
		}, nil, "failed to connect")
		return nil, err
	}
	return tx, nil
}

func InitDB(varenv models.VarEnviroment, log *logging.Logger) *DatabaseData {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", varenv.Host, varenv.User, varenv.Pass, varenv.DB, varenv.Port)
	print(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &DatabaseData{
		DB:  db,
		log: log,
	}
}
