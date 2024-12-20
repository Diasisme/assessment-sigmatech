package payload

import "time"

type RegisterReq struct {
	AccountNumber string `json:"account_number" validate:"required"`
	Pin           string `json:"pin" validate:"required"`
}

type UserReq struct {
	Nik           string    `json:"nik" validate:"required"`
	FullName      string    `json:"full_name" validate:"required"`
	LegalName     string    `json:"legal_name" validate:"required"`
	Birthplace    string    `json:"birthplace" validate:"required"`
	BirthDate     time.Time `json:"birth_date" validate:"required"`
	Salary        float64   `json:"salary" validate:"required"`
	CardID        int64     `json:"card_id" validate:"required"`
	AccountNumber string    `json:"account_number" validate:"required"`
}

type ActivationAccountReq struct {
	ID            int64  `json:"id" validate:"required"`
	AccountNumber string `json:"account_number" validate:"required"`
}

type UploadReq struct {
	ID int64 `json:"id" validate:"required"`
}

type TransactionReq struct {
	AccountNumber string  `json:"account_number" validate:"required"`
	UserID        int64   `json:"user_id" validate:"required"`
	LoanID        int64   `json:"loan_id" validate:"required"`
	Otr           float64 `json:"otr" validate:"required"`
	AssetName     string  `json:"asset_name" validate:"required"`
}

type TransferReq struct {
	NomorRekeningAsal   string  `json:"nomor_rekening_asal" validate:"required"`
	NomorRekeningTujuan string  `json:"nomor_rekening_tujuan" validate:"required"`
	Nominal             float64 `json:"nominal" validate:"required"`
}

type GetTransaksiReq struct {
	NomorRekening string `json:"nomor_rekening"`
}
