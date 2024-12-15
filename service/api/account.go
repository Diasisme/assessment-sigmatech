package api

import (
	"assesment-sigmatech/service/helpers"
	"assesment-sigmatech/service/models"
	"assesment-sigmatech/service/payload"
	"net/http"
	"strconv"
	"time"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (f *AccountApi) CreateAccount(c echo.Context) (err error) {

	startTime := time.Now()
	var request models.Account
	var response helpers.Response

	f.log.Info(logrus.Fields{
		"request": request,
	}, request, "log info")

	payloadValidator := new(payload.UserReq)

	if err = c.Bind(payloadValidator); err != nil {
		remark := "Cannot process bind validator. Please try again"
		f.log.Error(logrus.Fields{
			"error": err,
			"data":  payloadValidator,
		}, nil, remark)
		response.Message = remark
		response.Status = http.StatusBadRequest
		response.Data = nil

		err = c.JSON(response.Status, response)
		return
	}

	if err = f.validate.Struct(payloadValidator); err != nil {
		remark := "Data is not valid/empty. Please try again"
		f.log.Error(logrus.Fields{
			"error": err,
			"data":  payloadValidator,
		}, nil, remark)
		response.Message = remark
		response.Status = http.StatusBadRequest
		response.Data = nil

		err = c.JSON(response.Status, response)
		return
	}

	err = copier.Copy(&request, payloadValidator)
	if err != nil {
		remark := "Cannot process copy data from validator. Please try again."
		f.log.Error(logrus.Fields{
			"error":            err,
			"source copy":      payloadValidator,
			"destination copy": request,
		}, nil, remark)
		response.Message = remark
		response.Status = http.StatusBadRequest
		response.Data = nil

		err = c.JSON(response.Status, response)
		return
	}

	result, err := f.app.CreateAccount(request)
	if err != nil {
		remark := result.Message
		f.log.Error(logrus.Fields{
			"error": err.Error(),
		}, nil, remark)

		response.Message = result.Message
		response.Status = result.Status
		response.Data = nil

		err = c.JSON(response.Status, response)
		return
	}

	elapsedTime := time.Since(startTime)
	f.log.Info(logrus.Fields{
		"Request":     request,
		"Result":      result,
		"error":       err,
		"elapsedTime": elapsedTime,
	}, nil, "log info")

	return c.JSON(http.StatusOK, result)
}

func (f *AccountApi) UploadIDPhoto(c echo.Context) (err error) {
	startTime := time.Now()
	var request models.Account
	var response helpers.Response

	id := c.FormValue("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID is required"})
	}

	AccountID, _ := strconv.Atoi(id)

	f.log.Info(logrus.Fields{
		"request": request,
	}, request, "log info")

	result, err := f.app.UploadIDPhoto(c, int64(AccountID))
	if err != nil {
		remark := result.Message
		f.log.Error(logrus.Fields{
			"error": err.Error(),
		}, nil, remark)

		response.Message = result.Message
		response.Status = result.Status
		response.Data = nil

		err = c.JSON(response.Status, response)
		return
	}

	elapsedTime := time.Since(startTime)
	f.log.Info(logrus.Fields{
		"Request":     request,
		"Result":      result,
		"error":       err,
		"elapsedTime": elapsedTime,
	}, nil, "log info")

	return c.JSON(http.StatusOK, result)
}

func (f *AccountApi) UploadSelfiePhoto(c echo.Context) (err error) {
	startTime := time.Now()
	var request models.Account
	var response helpers.Response

	id := c.FormValue("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID is required"})
	}

	AccountID, _ := strconv.Atoi(id)

	f.log.Info(logrus.Fields{
		"request": request,
	}, request, "log info")

	result, err := f.app.UploadSelfiePhoto(c, int64(AccountID))
	if err != nil {
		remark := result.Message
		f.log.Error(logrus.Fields{
			"error": err.Error(),
		}, nil, remark)

		response.Message = result.Message
		response.Status = result.Status
		response.Data = nil

		err = c.JSON(response.Status, response)
		return
	}

	elapsedTime := time.Since(startTime)
	f.log.Info(logrus.Fields{
		"Request":     request,
		"Result":      result,
		"error":       err,
		"elapsedTime": elapsedTime,
	}, nil, "log info")

	return c.JSON(http.StatusOK, result)
}
