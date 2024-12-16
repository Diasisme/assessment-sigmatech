package repository

import (
	"assesment-sigmatech/service/models"

	"gorm.io/gorm"
)

func (f *DatabaseData) GetDataCard(tx *gorm.DB, id int64) (models.Card, error) {
	var data models.Card
	result := tx.Where("id", id).First(&data)
	return data, result.Error
}
