package repository

import "assesment-sigmatech/service/models"

func (f *DatabaseData) Register(request models.UserLogin) (err error) {
	return f.DB.Create(&request).Error
}
