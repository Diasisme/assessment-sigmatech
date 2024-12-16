package app

import (
	"assesment-sigmatech/service/helpers"
	"assesment-sigmatech/service/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AccountDatastore interface {
	Begin() (*gorm.DB, error)
	Register(tx *gorm.DB, request models.UserLogin) error
	CreateAccount(tx *gorm.DB, request models.Account) error
	// TambahTabung(request models.Tabungan) error
	// KurangTabung(request models.Tabungan) error
	// Transaksi(request models.Transaksi) error
	GetDataAccount(tx *gorm.DB, id int64) (models.Account, error)
	UpdateIDPhoto(tx *gorm.DB, id int64, url_id string) error
	UpdateSelfiePhoto(tx *gorm.DB, id int64, url_id string) error
	UpdateStatusAccount(tx *gorm.DB, id int64) error
	// GetDataTabungan(nomor_rekening string) (models.Tabungan, error)
	// GetSaldoTabungan(nomor_rekening string) (models.Tabungan, error)
	// Mutasi(nomor_rekening string) ([]models.Transaksi, error)
}

type AccountApp interface {
	Register(request models.UserLogin) (response helpers.Response, err error)
	CreateAccount(request models.Account) (response helpers.Response, err error)
	UploadIDPhoto(c echo.Context, account_id int64) (response helpers.Response, err error)
	UploadSelfiePhoto(c echo.Context, account_id int64) (response helpers.Response, err error)
	AccountActivation(c echo.Context, request models.Account) (response helpers.Response, err error)
	// Tarik(request models.Tabungan) (response helpers.Response, err error)
}
