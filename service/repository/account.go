package repository

import (
	"assesment-sigmatech/service/models"
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (f *DatabaseData) CreateAccount(tx *gorm.DB, request models.Account) (err error) {
	return tx.Create(&request).Error
}

func (f *DatabaseData) GetDataAccount(tx *gorm.DB, id int64) (models.Account, error) {
	var data models.Account
	result := tx.Where("id", id).First(&data)
	return data, result.Error
}

func (f *DatabaseData) UpdateIDPhoto(tx *gorm.DB, id int64, url_id string) (err error) {

	log.Info(logrus.Fields{
		"id":     id,
		"url_id": url_id,
	}, nil, "request input")

	result := tx.Model(&models.Account{}).Where("id = ?", id).Update("id_photo", url_id)
	if result.Error != nil {
		return result.Error
	}

	log.Info(logrus.Fields{
		"affected_rows": result.RowsAffected,
		"error":         result.Error,
	}, nil, "update result")

	// Periksa apakah ada baris yang diupdate
	if result.RowsAffected == 0 {
		return fmt.Errorf("no account found with id %d", id)
	}

	return
}

func (f *DatabaseData) UpdateSelfiePhoto(tx *gorm.DB, id int64, url_id string) (err error) {

	log.Info(logrus.Fields{
		"id":     id,
		"url_id": url_id,
	}, nil, "request input")

	result := tx.Model(&models.Account{}).Where("id = ?", id).Update("selfie_photo", url_id)
	if result.Error != nil {
		return result.Error
	}

	log.Info(logrus.Fields{
		"affected_rows": result.RowsAffected,
		"error":         result.Error,
	}, nil, "update result")

	if result.RowsAffected == 0 {
		return fmt.Errorf("no account found with id %d", id)
	}

	return
}
