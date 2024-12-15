package repository

import (
	"assesment-sigmatech/service/models"

	"gorm.io/gorm"
)

func (f *DatabaseData) Register(tx *gorm.DB, request models.UserLogin) (err error) {
	return tx.Create(&request).Error
}
