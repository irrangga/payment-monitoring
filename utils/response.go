package utils

import "github.com/gin-gonic/gin"

type Responses struct {
	StatusCode int         `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type ErrorResponse struct {
	StatusCode int         `json:"status"`
	Error      interface{} `json:"error"`
}

func APIResponse(c *gin.Context, Message string, StatusCode int, Data interface{}) {

	jsonResponse := Responses{
		StatusCode: StatusCode,
		Message:    Message,
		Data:       Data,
	}

	if StatusCode >= 400 {
		c.JSON(StatusCode, jsonResponse)
		defer c.AbortWithStatus(StatusCode)
	} else {
		c.JSON(StatusCode, jsonResponse)
	}
}

func ValidatorErrorResponse(c *gin.Context, StatusCode int, Method string, Error interface{}) {
	errResponse := ErrorResponse{
		StatusCode: StatusCode,
		Error:      Error,
	}

	c.JSON(StatusCode, errResponse)
	defer c.AbortWithStatus(StatusCode)
}
