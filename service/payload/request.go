package payload

type RegisterReq struct {
	AccountNumber string `json:"account_number" validate:"required"`
	Pin           string `json:"pin" validate:"required"`
}

type TabunganReq struct {
	NomorRekening string  `json:"nomor_rekening" validate:"required"`
	Nominal       float64 `json:"nominal" validate:"required"`
}

type TransferReq struct {
	NomorRekeningAsal   string  `json:"nomor_rekening_asal" validate:"required"`
	NomorRekeningTujuan string  `json:"nomor_rekening_tujuan" validate:"required"`
	Nominal             float64 `json:"nominal" validate:"required"`
}

type GetTransaksiReq struct {
	NomorRekening string `json:"nomor_rekening"`
}
