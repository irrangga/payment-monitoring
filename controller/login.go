package controller

import (
	"net/http"
	"payment-monitoring/models"
	"payment-monitoring/usecase"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	login usecase.LoginUsecaseInterface
}

func NewLoginControllerImpl(r *gin.RouterGroup, login usecase.LoginUsecaseInterface) {
	handler := LoginController{login}

	r.POST("/login", handler.Login)
}

func (lc LoginController) Login(c *gin.Context) {
	var user models.OfficerAccount

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	var authD models.Auth

	authRes, err := lc.login.CreateAuth(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	authD.AuthUUID = authRes.AuthUUID
	authD.Username = authRes.Username
	authD.RoleID = authRes.RoleID
	authD.RoleName = authRes.RoleName

	var jwt models.JwtModel

	in, err := lc.login.SignIn(authD)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	jwt.OfficerID = authRes.OfficerID
	jwt.CustomerID = authRes.CustomerID
	jwt.Token = in
	jwt.Name = authRes.Name

	response := models.ResponseCustom{
		Status:  200,
		Message: "success",
		Data:    jwt,
	}
	c.JSON(http.StatusOK, response)
}
