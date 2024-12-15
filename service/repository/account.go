package repository

import "assesment-sigmatech/service/models"

func (f *DatabaseData) CreateAccount(request models.Account) (err error) {
	return f.DB.Create(&request).Error
}
