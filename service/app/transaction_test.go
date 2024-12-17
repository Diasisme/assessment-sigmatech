package app_test

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"

	"assesment-sigmatech/service/helpers"
	"assesment-sigmatech/service/models"
)

type mockRepo struct {
	mock.Mock
}

func (m *mockRepo) Begin() (*gorm.DB, error) {
	args := m.Called()
	return args.Get(0).(*gorm.DB), args.Error(1)
}

func (m *mockRepo) GetDataAccount(tx *gorm.DB, userID int64) (models.Account, error) {
	args := m.Called(tx, userID)
	return args.Get(0).(models.Account), args.Error(1)
}

func (m *mockRepo) GetDataCard(tx *gorm.DB, cardID int64) (models.Card, error) {
	args := m.Called(tx, cardID)
	return args.Get(0).(models.Card), args.Error(1)
}

func (m *mockRepo) GetDataLimitLoan(tx *gorm.DB, loanID int64) (models.LimitLoan, error) {
	args := m.Called(tx, loanID)
	return args.Get(0).(models.LimitLoan), args.Error(1)
}

func (m *mockRepo) GetTransactionHistData(tx *gorm.DB, accountNumber string) (models.TransactionHist, error) {
	args := m.Called(tx, accountNumber)
	return args.Get(0).(models.TransactionHist), args.Error(1)
}

func (m *mockRepo) CreateTransaction(tx *gorm.DB, transaction models.Transaction) error {
	args := m.Called(tx, transaction)
	return args.Error(0)
}

func (m *mockRepo) CreateTransactionHist(tx *gorm.DB, hist models.TransactionHist) error {
	args := m.Called(tx, hist)
	return args.Error(0)
}

func (m *mockRepo) UpdateTransactionHist(tx *gorm.DB, hist models.TransactionHist) error {
	args := m.Called(tx, hist)
	return args.Error(0)
}

type mockTx struct {
	mock.Mock
}

func (m *mockTx) Rollback() *gorm.DB {
	m.Called()
	return &gorm.DB{} // Return objek gorm.DB kosong sebagai placeholder
}

func (m *mockTx) Commit() *gorm.DB {
	m.Called()
	return &gorm.DB{}
}


type accountApp struct {
	accRepo *mockRepo
	log     interface{} // Logger dapat disesuaikan jika diperlukan
}

func (a *accountApp) CreateTransaction(c interface{}, request models.Transaction) (helpers.Response, error) {
	response := helpers.Response{}
	tx, err := a.accRepo.Begin()
	if err != nil {
		return response, err
	}
	defer func() {
		if r := recover(); r != nil || err != nil {
			tx.Rollback()
		}
	}()

	account, err := a.accRepo.GetDataAccount(tx, request.UserID)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Message = "Cannot fetch data. Please try again."
		return response, err
	}

	card, err := a.accRepo.GetDataCard(tx, account.CardID)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Message = "Cannot fetch card data. Please try again."
		return response, err
	}

	limitLoan, err := a.accRepo.GetDataLimitLoan(tx, request.LoanID)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Message = "Cannot fetch loan limit data. Please try again."
		return response, err
	}

	fmt.Print(card)
	fmt.Print(limitLoan)

	response.Status = http.StatusOK
	response.Message = "Transaction Success."
	return response, nil
}

func TestCreateTransaction_Success(t *testing.T) {
	mockRepo := new(mockRepo)
	mockTx := &gorm.DB{}

	app := &accountApp{
		accRepo: mockRepo,
	}

	request := models.Transaction{
		UserID: 1,
		Otr:    50000,
		LoanID: 1,
	}

	account := models.Account{
		ID:            1,
		CardNumber:    "1234567890",
		AccountStatus: 1,
		AccountNumber: "9876543210",
		CardID:        1,
	}
	card := models.Card{
		CardLimit: 100000,
	}
	LimitLoan := models.LimitLoan{
		CardID:        1,
		LimitValue:    100000,
		InterestValue: 0.05,
		Tenor:         12,
	}

	mockRepo.On("Begin").Return(mockTx, nil)
	mockRepo.On("GetDataAccount", mockTx, request.UserID).Return(account, nil)
	mockRepo.On("GetDataCard", mockTx, account.CardID).Return(card, nil)
	mockRepo.On("GetDataLimitLoan", mockTx, request.LoanID).Return(LimitLoan, nil)

	response, err := app.CreateTransaction(nil, request)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.Status)
	assert.Equal(t, "Transaction Success.", response.Message)
	mockRepo.AssertExpectations(t)
}

// Still Error because return nil pointer at Rollback
func TestCreateTransaction_Rollback(t *testing.T) {
	mockRepo := new(mockRepo)
	mockTx := new(mockTx) // Gunakan mockTx sebagai transaction

	// Tambahkan expect pada Rollback di mockTx
	mockTx.On("Rollback").Return(&gorm.DB{})

	app := &accountApp{
		accRepo: mockRepo,
	}

	request := models.Transaction{
		UserID: 1,
		Otr:    50000,
		LoanID: 1,
	}

	mockRepo.On("Begin").Return(mockTx, nil)
	mockRepo.On("GetDataAccount", mockTx, request.UserID).Return(models.Account{}, errors.New("failed to fetch data"))

	response, err := app.CreateTransaction(nil, request)

	// Assert error dan rollback dipanggil
	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, response.Status)
	mockRepo.AssertCalled(t, "Begin")
	mockTx.AssertCalled(t, "Rollback") // Pastikan Rollback dipanggil
}
