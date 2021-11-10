package repository

import (
	"payment-monitoring/models"
	"time"

	"gorm.io/gorm"
)

type PaymentRepoStruct struct {
	db *gorm.DB
}

func NewPaymentRepoImpl(db *gorm.DB) PaymentRepoInterface {
	return PaymentRepoStruct{db}
}

func (pr PaymentRepoStruct) CreatePaymentRepository(input models.EntityPayment) (models.EntityPayment, error) {
	var payment models.EntityPayment
	db := pr.db.Model(&payment)

	layout := "2006-1-02 15:04:05"
	tm := time.Now()
	res := tm.Format(layout)

	payment.UserID = input.UserID
	payment.RequestedBy = input.RequestedBy
	payment.PaymentDetails = input.PaymentDetails
	payment.PaymentDate = input.PaymentDate
	payment.PaymentAmount = input.PaymentAmount
	payment.AmountInWords = input.AmountInWords
	payment.BeneficiaryAccountName = input.BeneficiaryAccountName
	payment.BeneficiaryAccountNumber = input.BeneficiaryAccountNumber
	payment.StatusRequest = 1
	payment.RequestSent = res

	err := db.Debug().Create(&payment).Error
	db.Commit()

	return payment, err
}

func (pr PaymentRepoStruct) GetReqByIdUnitRepository(id string) ([]models.EntityPayment, error) {
	var payment []models.EntityPayment
	err := pr.db.Where("user_id = ?", id).Find(&payment).Error
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (pr PaymentRepoStruct) GetReqByIdPaymentRepository(id int) (models.EntityPayment, error) {
	var payment models.EntityPayment
	err := pr.db.Debug().Where("payment_id = ?", id).First(&payment).Error
	if err != nil {
		return models.EntityPayment{}, err
	}
	return payment, nil
}

func (pr PaymentRepoStruct) UpdateAprovalRepository(approval models.EntityPayment) (models.EntityPayment, error) {
	trx := pr.db.Begin()

	err := pr.db.Debug().Save(&approval).Error
	if err != nil {
		trx.Rollback()
		return models.EntityPayment{}, err
	}

	trx.Commit()
	return approval, nil
}
