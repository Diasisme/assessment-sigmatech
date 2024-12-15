package app

import (
	"assesment-sigmatech/config/logging/utils"
	"assesment-sigmatech/service/helpers"
	"assesment-sigmatech/service/models"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (f *accountApp) CreateAccount(request models.Account) (response helpers.Response, err error) {

	request.IDPhoto = ""
	request.SelfiePhoto = ""
	request.AccountStatus = 0

	cardNumber := utils.GenerateRandomNumber(16)
	request.CardNumber = strconv.Itoa(cardNumber)

	if err = f.accRepo.CreateAccount(request); err != nil {
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
		"card_number":    request.CardNumber,
		"account_number": request.AccountNumber,
	}

	return
}
