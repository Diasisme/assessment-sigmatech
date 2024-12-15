package app

import (
	"assesment-sigmatech/config/logging/utils"
	"assesment-sigmatech/service/helpers"
	"assesment-sigmatech/service/models"
	"net/http"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

func (f *accountApp) Register(request models.UserLogin) (response helpers.Response, err error) {

	tx, err := f.accRepo.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
		if err != nil {
			tx.Rollback()
		}
	}()

	hashPin, _ := utils.HashPin(request.Pin)
	request.Pin = hashPin

	if err = f.accRepo.Register(tx, request); err != nil {
		remark := "Failed to enter data. Please try again."
		f.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = http.StatusBadRequest
		response.Message = remark
		err = status.Error(http.StatusBadRequest, err.Error())
		return
	}

	if err = tx.Commit().Error; err != nil {
		remark := "Transaction failed. Please try again."
		f.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = http.StatusInternalServerError
		response.Message = remark
		return
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = utils.JSON{
		"account_number": request.AccountNumber,
	}

	return
}
