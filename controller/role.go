package controller

import (
	"net/http"
	"payment-monitoring/models"
	"payment-monitoring/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	role usecase.RoleUsecaseInterface
}

func NewRoleControllerImpl(r *gin.RouterGroup, role usecase.RoleUsecaseInterface) {
	handler := RoleController{role}

	r.POST("/role/create", handler.AddRoleController)
	r.GET("/list_role", handler.GetAllRoleController)
	r.GET("/role/:id", handler.GetRoleController)
	r.PUT("/role/:id", handler.UpdateRoleController)
	r.DELETE("/role/:id", handler.DeleteRoleController)
}

func (rc RoleController) AddRoleController(c *gin.Context) {
	var role models.Role

	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	roleData, err := rc.role.AddRoleUsecase(role)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	response := models.ResponseCustom{
		Status:  http.StatusOK,
		Message: "success",
		Data:    roleData,
	}
	c.JSON(http.StatusOK, response)
}

func (rc RoleController) GetAllRoleController(c *gin.Context) {
	roleData, err := rc.role.GetAllRoleUsecase()
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	response := models.ResponseCustom{
		Status:  http.StatusOK,
		Message: "success",
		Data:    roleData,
	}
	c.JSON(http.StatusOK, response)
}

func (rc RoleController) GetRoleController(c *gin.Context) {
	var role models.Role
	id := c.Param("id")

	roleData := rc.role.GetRoleUsecase(role, id)
	response := models.ResponseCustom{
		Status:  http.StatusOK,
		Message: "success",
		Data:    roleData,
	}
	c.JSON(http.StatusOK, response)
}

func (rc RoleController) UpdateRoleController(c *gin.Context) {
	var role models.Role
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	err := c.ShouldBindJSON(&role)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	roleData, err := rc.role.UpdateRoleUsecase(role, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	roleData.ID = idInt

	response := models.ResponseCustom{
		Status:  http.StatusOK,
		Message: "success",
		Data:    roleData,
	}
	c.JSON(http.StatusOK, response)
}

func (rc RoleController) DeleteRoleController(c *gin.Context) {
	var role models.Role
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	rc.role.DeleteRoleUsecase(role, id)

	accountID := &models.RoleID{
		ID: idInt,
	}

	response := models.ResponseCustom{
		Status:  http.StatusOK,
		Message: "success",
		Data:    accountID,
	}
	c.JSON(http.StatusOK, response)
}
