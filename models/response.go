package models

type ResponseCustom struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseErrorCustom struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
}
