package api

import (
	"assesment-sigmatech/service/helpers"
	"assesment-sigmatech/service/models"
	"assesment-sigmatech/service/payload"
	"net/http"
	"time"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (f *AccountApi) Register(c echo.Context) (err error) {
	startTime := time.Now()
	var request models.UserLogin
	var response helpers.Response

	f.log.Info(logrus.Fields{
		"request": request,
	}, request, "log info")

	payloadValidator := new(payload.RegisterReq)

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
		return err
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
		return err
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
		return err
	}

	result, err := f.app.Register(request)
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
