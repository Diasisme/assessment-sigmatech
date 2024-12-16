package repository

import (
	"assesment-sigmatech/service/models"

	"gorm.io/gorm"
)

func (f *DatabaseData) GetDataLimitLoan(tx *gorm.DB, id int64) (*models.LimitLoan, error) {
	var data models.LimitLoan
	result := tx.Where("id", id).First(&data)
	return &data, result.Error
}
