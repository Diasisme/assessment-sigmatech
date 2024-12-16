package repository

import (
	"assesment-sigmatech/service/models"
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (f *DatabaseData) GetTransactionHistData(tx *gorm.DB, account_number string) (*models.TransactionHist, error) {
	var data models.TransactionHist
	result := tx.Where("account_number", account_number).First(&data)
	return &data, result.Error
}

func (f *DatabaseData) CreateTransaction(tx *gorm.DB, request models.Transaction) (err error) {
	return tx.Create(&request).Error
}

func (f *DatabaseData) CreateTransactionHist(tx *gorm.DB, request models.TransactionHist) (err error) {
	return tx.Create(&request).Error
}

func (f *DatabaseData) UpdateTransactionHist(tx *gorm.DB, request models.TransactionHist) (err error) {
	log.Info(logrus.Fields{
		"id":         request.ID,
		"total_loan": request.TotalLoan,
	}, nil, "request input")

	result := tx.Model(&models.TransactionHist{}).Where("id = ?", request.ID).Update("total_loan", request.TotalLoan)
	if result.Error != nil {
		return result.Error
	}

	log.Info(logrus.Fields{
		"affected_rows": result.RowsAffected,
		"error":         result.Error,
	}, nil, "update result")

	if result.RowsAffected == 0 {
		return fmt.Errorf("no transaction history found with id %d", request.ID)
	}

	return
}
