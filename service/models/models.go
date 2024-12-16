package models

import (
	"time"
)

// Struct UserLogin (credit.user_login)
type UserLogin struct {
	AccountNumber string `gorm:"column:account_number;primaryKey"`
	Pin           string `gorm:"column:pin;not null"`
}

// Struct Card (credit.card)
type Card struct {
	ID        int64   `gorm:"column:id;primaryKey;autoIncrement"`
	TierCard  string  `gorm:"column:tier_card;not null;unique"`
	CardLimit float64 `gorm:"column:card_limit;not null"`
}

// Struct User (credit.user)
type Account struct {
	ID            int64     `gorm:"column:id;primaryKey;autoIncrement"`
	Nik           string    `gorm:"column:nik;not null;unique"`
	FullName      string    `gorm:"column:full_name;not null"`
	LegalName     string    `gorm:"column:legal_name;not null"`
	Birthplace    string    `gorm:"column:birthplace;not null"`
	BirthDate     time.Time `gorm:"column:birth_date;not null"`
	Salary        float64   `gorm:"column:salary;not null"`
	IDPhoto       string    `gorm:"column:id_photo;not null"`
	SelfiePhoto   string    `gorm:"column:selfie_photo;not null"`
	CardNumber    string    `gorm:"column:card_number;not null;unique"`
	AccountStatus int64     `gorm:"column:user_status;not null"`
	CardID        int64     `gorm:"column:card_id"`
	AccountNumber string    `gorm:"column:account_number"`

	Card      Card      `gorm:"foreignKey:CardID;references:ID"`
	UserLogin UserLogin `gorm:"foreignKey:AccountNumber;references:AccountNumber"`
}

// Struct LimitLoan (credit.limit_loan)
type LimitLoan struct {
	ID            int64   `gorm:"column:id;primaryKey;autoIncrement"`
	CardID        int64   `gorm:"column:card_id"`
	Tenor         int     `gorm:"column:tenor;not null"`
	LimitValue    float64 `gorm:"column:limit_value;not null"`
	InterestValue float64 `gorm:"column:interest_value;not null"`

	Card Card `gorm:"foreignKey:CardID;references:ID"`
}

// Struct Transaction (credit.transaction)
type Transaction struct {
	ID                int64     `gorm:"column:id;primaryKey;autoIncrement"`
	UserID            int64     `gorm:"column:user_id"`
	ContractNo        string    `gorm:"column:contract_no;not null;unique"`
	LoanID            int64     `gorm:"column:loan_id"`
	Otr               float64   `gorm:"column:otr;not null"`
	AdminFee          float64   `gorm:"column:admin_fee;not null"`
	InstallmentValue  float64   `gorm:"column:installment_value;not null"`
	InterestAmount    float64   `gorm:"column:interest_amount;not null"`
	AssetName         string    `gorm:"column:asset_name;not null"`
	AccountNumber     string    `gorm:"column:account_number"`
	CardNumber        string    `gorm:"column:card_number;not null"`
	TransactionStatus bool      `gorm:"column:transaction_status;not null;default:true"`
	InputDate         time.Time `gorm:"column:input_date;not null"`

	Loan      LimitLoan `gorm:"foreignKey:LoanID;references:ID"`
	User      Account   `gorm:"foreignKey:UserID;references:ID"`
	UserLogin UserLogin `gorm:"foreignKey:AccountNumber;references:AccountNumber"`
}

type TransactionHist struct {
	ID            int64     `gorm:"column:id;primaryKey;autoIncrement"`
	AccountNumber string    `gorm:"column:account_number"`
	UserID        int64     `gorm:"column:user_id"`
	TotalLoan     float64   `gorm:"column:total_loan;not null"`
	InputDate     time.Time `gorm:"column:input_date;not null"`
	UpdateDate    time.Time `gorm:"column:update_date"`

	User      Account   `gorm:"foreignKey:UserID;references:ID"`
	UserLogin UserLogin `gorm:"foreignKey:AccountNumber;references:AccountNumber"`
}

type FileMetadata struct {
	FileName string `json:"file_name"`
	FileSize int64  `json:"file_size"`
	FileURL  string `json:"file_url"`
}

type UploadToCloudResult struct {
	Path   string
	Bucket string
	Url    string
}

type FileRecord struct {
	ID       string `gorm:"id"`
	FilePath string
}

func (UserLogin) TableName() string {
	return "credit.user_login"
}

func (Card) TableName() string {
	return "credit.card"
}

func (Account) TableName() string {
	return "credit.user"
}

func (LimitLoan) TableName() string {
	return "credit.limit_loan"
}

func (Transaction) TableName() string {
	return "credit.transaction"
}

func (TransactionHist) TableName() string {
	return "credit.transaction_hist"
}
