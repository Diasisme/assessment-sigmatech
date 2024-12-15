package app

import (
	"assesment-sigmatech/config/logging/utils"
	"assesment-sigmatech/service/helpers"
	"assesment-sigmatech/service/models"
	"net/http"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (f *accountApp) Register(request models.UserLogin) (response helpers.Response, err error) {

	hashPin, _ := utils.HashPin(request.Pin)
	request.Pin = hashPin

	if err = f.accRepo.Register(request); err != nil {
		remark := "Failed to enter data. Please try again."
		f.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = http.StatusBadRequest
		response.Message = remark
		err = status.Error(codes.InvalidArgument, err.Error())
		return
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = utils.JSON{
		"account_number": request.AccountNumber,
	}

	return
}
