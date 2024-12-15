package app

import (
	"assesment-sigmatech/config/logging/utils"
	"assesment-sigmatech/service/helpers"
	"assesment-sigmatech/service/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (f *accountApp) CreateAccount(request models.Account) (response helpers.Response, err error) {

	tx, err := f.accRepo.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
		if err != nil {
			tx.Rollback()
		}
	}()

	request.IDPhoto = ""
	request.SelfiePhoto = ""
	request.AccountStatus = 0

	cardNumber := utils.GenerateRandomNumber(16)
	request.CardNumber = strconv.Itoa(cardNumber)

	if err = f.accRepo.CreateAccount(tx, request); err != nil {
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
		"card_number":    request.CardNumber,
		"account_number": request.AccountNumber,
	}

	return
}

func (f *accountApp) UploadIDPhoto(c echo.Context, account_id int64) (response helpers.Response, err error) {

	tx, err := f.accRepo.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
		if err != nil {
			tx.Rollback()
		}
	}()

	getData, err := f.accRepo.GetDataAccount(tx, account_id)
	if err != nil && err != gorm.ErrRecordNotFound {
		remark := "Cannot fetch data. Please try again."
		f.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = http.StatusBadRequest
		response.Message = remark
		err = status.Error(http.StatusBadRequest, err.Error())
		return
	}

	filename := fmt.Sprintf("KTP_%[1]s_%[2]s.png", getData.Nik, getData.AccountNumber)

	resultMinio, err := f.minioClient.UploadToCloud(c, "ktp", filename, "PNG")
	if err != nil {
		remark := err.Error()
		f.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = http.StatusBadRequest
		response.Message = remark
		return
	}
	pathIDPhoto := resultMinio.Url

	if err = f.accRepo.UpdateIDPhoto(tx, account_id, pathIDPhoto); err != nil {
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
		"account_id":     account_id,
		"account_number": getData.AccountNumber,
	}

	return
}

func (f *accountApp) UploadSelfiePhoto(c echo.Context, account_id int64) (response helpers.Response, err error) {

	tx, err := f.accRepo.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
		if err != nil {
			tx.Rollback()
		}
	}()

	getData, err := f.accRepo.GetDataAccount(tx, account_id)
	if err != nil && err != gorm.ErrRecordNotFound {
		remark := "Cannot fetch data. Please try again."
		f.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = http.StatusBadRequest
		response.Message = remark
		err = status.Error(http.StatusBadRequest, err.Error())
		return
	}

	filename := fmt.Sprintf("SELFIE_%[1]s_%[2]s.png", getData.Nik, getData.AccountNumber)

	resultMinio, err := f.minioClient.UploadToCloud(c, "selfie", filename, "PNG")
	if err != nil {
		remark := err.Error()
		f.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = http.StatusBadRequest
		response.Message = remark
		return
	}
	pathSelfiePhoto := resultMinio.Url

	if err = f.accRepo.UpdateSelfiePhoto(tx, account_id, pathSelfiePhoto); err != nil {
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
		"account_id":     account_id,
		"account_number": getData.AccountNumber,
	}

	return
}
