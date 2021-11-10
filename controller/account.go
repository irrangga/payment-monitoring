package controller

import (
	"net/http"
	"payment-monitoring/models"
	"payment-monitoring/usecase"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
	account usecase.AccountUsecaseInterface
}

func NewAccountControllerImpl(r *gin.RouterGroup, account usecase.AccountUsecaseInterface) {
	handler := AccountController{account}

	r.POST("/account/create", handler.AddAccountController)
	r.GET("/list_account", handler.GetAllAccountController)
	r.GET("/account/:id", handler.GetAccountController)
	r.PUT("/account/:id", handler.UpdateAccountController)
	r.DELETE("/account/:id", handler.DeleteAccountController)
}

func (ac AccountController) AddAccountController(c *gin.Context) {
	var accOfficer models.OfficerAccount
	var userID models.User

	if errOfficer := c.ShouldBindJSON(&accOfficer); errOfficer != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: errOfficer.Error(),
		})
		return
	}

	if accOfficer.Role == 0 {
		accOfficer.Role = 4
	}

	if accOfficer.UserID == 0 {
		accUserID := ac.account.AddUserIdUsecase(userID)
		accOfficer.UserID = accUserID.ID
	}

	accDataOfficer, accDataCustomer, err := ac.account.AddAccountUsecase(accOfficer)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if accDataOfficer.Role == 0 {
		response := models.ResponseCustom{
			Status:  http.StatusOK,
			Message: "success",
			Data:    accDataCustomer,
		}
		c.JSON(http.StatusOK, response)
	} else {
		response := models.ResponseCustom{
			Status:  http.StatusOK,
			Message: "success",
			Data:    accDataOfficer,
		}
		c.JSON(http.StatusOK, response)
	}
}

func (ac AccountController) GetAllAccountController(c *gin.Context) {
	accDataOfficer, errOfficer := ac.account.GetAllOfficerAccountUsecase()
	accDataCustomer, errCustomer := ac.account.GetAllWorkUnitAccountUsecase()

	accDataOfficer = accDataOfficer[:len(accDataOfficer)+len(accDataCustomer)]

	for i := range accDataCustomer {
		accDataOfficer[len(accDataOfficer)-len(accDataCustomer)+i].UserID = accDataCustomer[i].UserID
		accDataOfficer[len(accDataOfficer)-len(accDataCustomer)+i].Name = accDataCustomer[i].Name
		accDataOfficer[len(accDataOfficer)-len(accDataCustomer)+i].Role = 4
		accDataOfficer[len(accDataOfficer)-len(accDataCustomer)+i].Username = accDataCustomer[i].Username
		accDataOfficer[len(accDataOfficer)-len(accDataCustomer)+i].Password = accDataCustomer[i].Password
	}

	sort.Slice(accDataOfficer, func(i, j int) bool {
		return accDataOfficer[i].UserID < accDataOfficer[j].UserID
	})

	if errOfficer != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: errOfficer.Error(),
		})
		return
	} else if errCustomer != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: errCustomer.Error(),
		})
		return
	}

	response := models.ResponseCustom{
		Status:  http.StatusOK,
		Message: "success",
		Data:    accDataOfficer,
	}

	c.JSON(http.StatusOK, response)
}

func (ac AccountController) GetAccountController(c *gin.Context) {
	var accOfficer models.OfficerAccount
	var accCustomer models.WorkUnitAccount
	id := c.Param("id")

	accDataOfficer, accDataCustomer := ac.account.GetAccountUsecase(accOfficer, accCustomer, id)
	if accDataOfficer.Role == 0 {
		response := models.ResponseCustom{
			Status:  http.StatusOK,
			Message: "success",
			Data:    accDataCustomer,
		}
		c.JSON(http.StatusOK, response)
	} else {
		response := models.ResponseCustom{
			Status:  http.StatusOK,
			Message: "success",
			Data:    accDataOfficer,
		}
		c.JSON(http.StatusOK, response)
	}
}

func (ac AccountController) UpdateAccountController(c *gin.Context) {
	var accOfficer models.OfficerAccount
	var accCustomer models.WorkUnitAccount
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	err := c.ShouldBindJSON(&accOfficer)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	accDataOfficer, accDataCustomer, err := ac.account.UpdateAccountUsecase(accOfficer, accCustomer, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	accDataOfficer.UserID, accDataCustomer.UserID = idInt, idInt

	if accDataOfficer.Role == 0 {
		response := models.ResponseCustom{
			Status:  http.StatusOK,
			Message: "success",
			Data:    accDataCustomer,
		}
		c.JSON(http.StatusOK, response)
	} else {
		response := models.ResponseCustom{
			Status:  http.StatusOK,
			Message: "success",
			Data:    accDataOfficer,
		}
		c.JSON(http.StatusOK, response)
	}
}

func (ac AccountController) DeleteAccountController(c *gin.Context) {
	var accOfficer models.OfficerAccount
	var accCustomer models.WorkUnitAccount
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	ac.account.DeleteAccountUsecase(accOfficer, accCustomer, id)

	accountID := &models.User{
		ID: idInt,
	}

	response := models.ResponseCustom{
		Status:  http.StatusOK,
		Message: "success",
		Data:    accountID,
	}
	c.JSON(http.StatusOK, response)
}
