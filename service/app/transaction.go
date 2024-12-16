package app

import (
	"assesment-sigmatech/config/logging/utils"
	"assesment-sigmatech/service/helpers"
	"assesment-sigmatech/service/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (f *accountApp) CreateTransaction(c echo.Context, request models.Transaction) (response helpers.Response, err error) {

	tx, err := f.accRepo.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
		if err != nil {
			tx.Rollback()
		}
	}()

	// Get Account Data
	getData, err := f.accRepo.GetDataAccount(tx, request.UserID)
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

	if err == gorm.ErrRecordNotFound {
		remark := "Data not found. Please try again."
		f.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = http.StatusBadRequest
		response.Message = remark
		err = status.Error(http.StatusBadRequest, err.Error())
		return
	}

	if getData.AccountStatus != 1 {
		remark := "Account not yet verified."
		f.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = http.StatusBadRequest
		response.Message = remark
		err = status.Error(http.StatusBadRequest, err.Error())
		return
	}

	request.InputDate = time.Now()
	request.CardNumber = getData.CardNumber
	request.AdminFee = 10000

	// Get Card Data
	getCardData, err := f.accRepo.GetDataCard(tx, getData.CardID)
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

	if err == gorm.ErrRecordNotFound {
		remark := "Data not found. Please try again."
		f.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = http.StatusBadRequest
		response.Message = remark
		err = status.Error(http.StatusBadRequest, err.Error())
		return
	}

	// Get Limit Loan Data
	getLimitLoan, err := f.accRepo.GetDataLimitLoan(tx, request.LoanID)
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

	if err == gorm.ErrRecordNotFound {
		remark := "Data not found. Please try again."
		f.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = http.StatusBadRequest
		response.Message = remark
		err = status.Error(http.StatusBadRequest, err.Error())
		return
	}

	if getData.CardID != getLimitLoan.CardID {
		remark := fmt.Sprintf("Card with ID %d is not valid. Please Try Again", getData.CardID)
		f.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = http.StatusBadRequest
		response.Message = remark
		err = status.Error(http.StatusBadRequest, err.Error())
		return
	}

	if request.Otr > getLimitLoan.LimitValue {
		remark := "Payment must not exceed loan limit"
		f.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = http.StatusBadRequest
		response.Message = remark
		err = status.Error(http.StatusBadRequest, err.Error())
		return
	}

	// Calculate Interest Amount and Installment Value
	request.InterestAmount = (request.Otr * getLimitLoan.InterestValue * float64(getLimitLoan.Tenor))
	request.InstallmentValue = ((request.Otr + request.InterestAmount) / float64(getLimitLoan.Tenor))
	request.ContractNo = strconv.Itoa(utils.GenerateRandomNumber(16))

	getTrxHistData, err := f.accRepo.GetTransactionHistData(tx, request.AccountNumber)
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

	if err == gorm.ErrRecordNotFound {
		if err = f.accRepo.CreateTransaction(tx, request); err != nil {
			remark := "Failed to enter data. Please try again."
			f.log.Warn(logrus.Fields{
				"err": err,
			}, nil, remark)
			response.Status = http.StatusBadRequest
			response.Message = remark
			err = status.Error(http.StatusBadRequest, err.Error())
			return
		}

		inputTransactionHist := models.TransactionHist{
			AccountNumber: request.AccountNumber,
			UserID:        request.UserID,
			TotalLoan:     request.InstallmentValue,
			InputDate:     time.Now(),
		}

		if err = f.accRepo.CreateTransactionHist(tx, inputTransactionHist); err != nil {
			remark := "Failed to enter data. Please try again."
			f.log.Warn(logrus.Fields{
				"err": err,
			}, nil, remark)
			response.Status = http.StatusBadRequest
			response.Message = remark
			err = status.Error(http.StatusBadRequest, err.Error())
			return
		}
	} else {
		loanValue := getTrxHistData.TotalLoan + request.InstallmentValue
		if loanValue > getCardData.CardLimit {
			remark := "Your loan is due limit, cannot request loan."
			f.log.Warn(logrus.Fields{
				"err": err,
			}, nil, remark)
			response.Status = http.StatusBadRequest
			response.Message = remark
			err = status.Error(http.StatusBadRequest, err.Error())
			return
		}

		if err = f.accRepo.CreateTransaction(tx, request); err != nil {
			remark := "Failed to enter data. Please try again."
			f.log.Warn(logrus.Fields{
				"err": err,
			}, nil, remark)
			response.Status = http.StatusBadRequest
			response.Message = remark
			err = status.Error(http.StatusBadRequest, err.Error())
			return
		}

		updateRequest := models.TransactionHist{
			TotalLoan:  loanValue,
			ID:         getTrxHistData.ID,
			UpdateDate: time.Now(),
		}

		if err = f.accRepo.UpdateTransactionHist(tx, updateRequest); err != nil {
			remark := "Failed to update data. Please try again."
			f.log.Warn(logrus.Fields{
				"err": err,
			}, nil, remark)
			response.Status = http.StatusBadRequest
			response.Message = remark
			err = status.Error(http.StatusBadRequest, err.Error())
			return
		}
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
	response.Message = "Transaction Success."
	response.Data = utils.JSON{
		"account_id":     getData.ID,
		"account_number": getData.AccountNumber,
	}

	return
}
