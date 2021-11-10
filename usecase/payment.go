package usecase

import (
	"payment-monitoring/models"
	"payment-monitoring/repository"
	"strconv"
)

type PaymentUsecaseStruct struct {
	payment repository.PaymentRepoInterface
}

func NewPaymentUsecaseImpl(payment repository.PaymentRepoInterface) PaymentUsecaseInterface {
	return &PaymentUsecaseStruct{payment}
}

func (ru PaymentUsecaseStruct) CreatePaymentUsecase(input models.InputCreatePayment) (models.EntityPayment, error) {

	payments := models.EntityPayment{
		UserID:                   input.UserID,
		RequestedBy:              input.RequestedBy,
		PaymentDetails:           input.PaymentDetails,
		PaymentDate:              input.PaymentDate,
		PaymentAmount:            input.PaymentAmount,
		AmountInWords:            input.AmountInWords,
		BeneficiaryAccountName:   input.BeneficiaryAccountName,
		BeneficiaryAccountNumber: input.BeneficiaryAccountNumber,
		StatusRequest:            1,
	}

	resultCreatePayment, errCreatePayment := ru.payment.CreatePaymentRepository(payments)

	return resultCreatePayment, errCreatePayment
}

func (ru PaymentUsecaseStruct) GetReqByIdUnitUsecase(id string) ([]models.EntityPayment, error) {
	ReqData, err := ru.payment.GetReqByIdUnitRepository(id)
	if err != nil {
		return nil, err
	}

	return ReqData, nil
}

func (ru PaymentUsecaseStruct) GetReqByIdPaymentUsecase(id string) (models.EntityPayment, error) {
	idRes, err := strconv.Atoi(id)
	if err != nil {
		return models.EntityPayment{}, err
	}
	payment, err := ru.payment.GetReqByIdPaymentRepository(idRes)
	if err != nil {
		return models.EntityPayment{}, err
	}

	return payment, nil
}

func (ru PaymentUsecaseStruct) UpdateAprovalUsecase(aproval models.EntityPayment, tipeReqRes string) (models.EntityPayment, error) {

	payment, err := ru.payment.GetReqByIdPaymentRepository(int(aproval.PaymentID))
	if err != nil {
		return models.EntityPayment{}, err
	}

	aproval.PaymentID = payment.PaymentID
	aproval.UserID = payment.UserID
	aproval.RequestedBy = payment.RequestedBy
	aproval.PaymentDetails = payment.PaymentDetails
	aproval.PaymentDate = payment.PaymentDate
	aproval.PaymentAmount = payment.PaymentAmount
	aproval.AmountInWords = payment.AmountInWords
	aproval.BeneficiaryAccountName = payment.BeneficiaryAccountName
	aproval.BeneficiaryAccountNumber = payment.BeneficiaryAccountNumber
	aproval.RequestSent = payment.RequestSent

	if tipeReqRes == "declined_by_gs" {
		aproval.StatusRequest = 2
	} else if tipeReqRes == "proceeding" {
		aproval.StatusRequest = 3
	} else if tipeReqRes == "declined_by_ac" {
		aproval.StatusRequest = 4
	} else if tipeReqRes == "approved" {
		aproval.StatusRequest = 5
	}

	repo, err := ru.payment.UpdateAprovalRepository(aproval)
	if err != nil {
		return models.EntityPayment{}, err
	}

	return repo, nil
}
