package models

import (
	"time"

	"gorm.io/gorm"
)

type Approval struct {
	gorm.Model
	PaymentID     int       `json:"payment_id"`
	StatusRequest int       `json:"status_request"`
	Reason        string    `json:"reason"`
	CreatedBy     string    `json:"created_by"`
	CreatedAt     time.Time `json:"created_at"`
}

type StatusRequest struct {
	Status     int    `json:"status"`
	StatusName string `json:"status_name"`
}

type EntityPayment struct {
	PaymentID                uint      `gorm:"primarykey;auto_increment;not null" json:"payment_id"`
	UserID                   uint      `gorm:"type:int(10);auto_increment; not null;" json:"user_id"`
	RequestedBy              string    `gorm:"type:varchar(255);not null;" json:"requested_by"`
	PaymentDetails           string    `gorm:"type:varchar(255);not null;" json:"payment_details"`
	PaymentDate              string    `json:"payment_date"`
	PaymentAmount            int64     `gorm:"not null;" json:"payment_amount"`
	AmountInWords            string    `gorm:"type:varchar(255);not null;" json:"amount_in_words"`
	BeneficiaryAccountName   string    `gorm:"type:varchar(255);not null;" json:"beneficiary_account_name"`
	BeneficiaryAccountNumber string    `gorm:"type:varchar(255);not null;" json:"beneficiary_account_number"`
	CreatedAt                time.Time `json:"-"`
	UpdatedAt                time.Time `json:"-"`
	RequestSent              string    `json:"request_sent"`
	StatusRequest            int       `json:"status"`
	Reason                   string    `json:"reason,omitempty"`
}

type InputCreatePayment struct {
	UserID                   uint   `json:"user_id" validate:"required,numeric"`
	RequestedBy              string `json:"requested_by" validate:"required"`
	PaymentDetails           string `json:"payment_details" validate:"required"`
	PaymentDate              string `json:"payment_date" validate:"required"`
	PaymentAmount            int64  `json:"payment_amount" validate:"required,numeric"`
	AmountInWords            string `json:"amount_in_words" validate:"required"`
	BeneficiaryAccountName   string `json:"beneficiary_account_name" validate:"required"`
	BeneficiaryAccountNumber string `json:"beneficiary_account_number" validate:"required"`
	RequestSent              string `json:"request_sent"`
	Status                   int    `json:"status"`
}
