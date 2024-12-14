package app

import (
	"assesment-sigmatech/service/helpers"
	"assesment-sigmatech/service/models"
)

type AccountDatastore interface {
	Register(request models.UserLogin) error
	// BuatTabung(request models.Tabungan) error
	// TambahTabung(request models.Tabungan) error
	// KurangTabung(request models.Tabungan) error
	// Transaksi(request models.Transaksi) error
	// GetDataAccount(nomor_rekening string) (models.Nasabah, error)
	// GetDataTabungan(nomor_rekening string) (models.Tabungan, error)
	// GetSaldoTabungan(nomor_rekening string) (models.Tabungan, error)
	// Mutasi(nomor_rekening string) ([]models.Transaksi, error)
}

type AccountApp interface {
	Register(request models.UserLogin) (response helpers.Response, err error)
	// Tabung(request models.Tabungan) (response helpers.Response, err error)
	// Tarik(request models.Tabungan) (response helpers.Response, err error)
}
