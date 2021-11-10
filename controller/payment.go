package controller

import (
	"net/http"
	"payment-monitoring/models"
	"payment-monitoring/usecase"
	"payment-monitoring/utils"
	"strings"

	"github.com/gin-gonic/gin"
	gpc "github.com/restuwahyu13/go-playground-converter"
)

type PaymentController struct {
	payment usecase.PaymentUsecaseInterface
}

func NewPaymentControllerImpl(r *gin.RouterGroup, payment usecase.PaymentUsecaseInterface) {
	handler := PaymentController{payment}

	r.POST("/payment/create", handler.CreatePaymentController)
	r.GET("/list_payment", handler.GetReqByIdUnitController)
	r.GET("/payment/:id", handler.GetReqByIdPaymentController)
	r.POST("/payment/status/:type", handler.ApprovalReqPayment)
}

func (pc PaymentController) CreatePaymentController(c *gin.Context) {
	var input models.InputCreatePayment
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	config := gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "UserID",
				Message: "user_id is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "numeric",
				Field:   "UserID",
				Message: "user_id must be numeric",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "RequestedBy",
				Message: "name is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "uppercase",
				Field:   "RequestedBy",
				Message: "requested_by must be using uppercase",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "PaymentDetails",
				Message: "payment_details is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "PaymentDate",
				Message: "payment_date is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "PaymentAmount",
				Message: "payment_amount is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "numeric",
				Field:   "PaymentAmount",
				Message: "payment_amount must be using number",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "AmountInWords",
				Message: "amount_in_words is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "BeneficiaryAccountName",
				Message: "beneficiary_account_name is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "BeneficiaryAccountNumber",
				Message: "beneficiary_account_number is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "numeric",
				Field:   "BeneficiaryAccountNumber",
				Message: "beneficiary_account_number is required on body",
			},
		},
	}

	errResponse, errCount := utils.GoValidator(&input, config.Options)

	if errCount > 0 {
		utils.ValidatorErrorResponse(c, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	result, errCreatePayment := pc.payment.CreatePaymentUsecase(input)

	switch errCreatePayment {

	default:
		utils.APIResponse(c, "success", http.StatusOK, result)
	}
}

func (pc PaymentController) GetReqByIdUnitController(c *gin.Context) {
	id := c.Query("user_id")

	reqData, err := pc.payment.GetReqByIdUnitUsecase(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	response := models.ResponseCustom{
		Status:  200,
		Message: "success",
		Data:    reqData,
	}
	c.JSON(http.StatusOK, response)
}

func (pc PaymentController) GetReqByIdPaymentController(c *gin.Context) {
	id := c.Param("id")

	payment, err := pc.payment.GetReqByIdPaymentUsecase(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	response := models.ResponseCustom{
		Status:  200,
		Message: "success",
		Data:    payment,
	}
	c.JSON(http.StatusOK, response)
}

func (pc PaymentController) ApprovalReqPayment(c *gin.Context) {
	var request models.EntityPayment

	err := c.ShouldBindJSON(&request)
	if err != nil {
		return
	}

	tipeReqTemp := c.Param("type")

	tipeReqRes := strings.ToLower(tipeReqTemp)

	approval, err := pc.payment.UpdateAprovalUsecase(request, tipeReqRes)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	response := models.ResponseCustom{
		Status:  200,
		Message: "success",
		Data:    approval,
	}
	c.JSON(http.StatusOK, response)
}
