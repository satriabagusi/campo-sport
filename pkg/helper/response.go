package helper

import "github.com/gin-gonic/gin"

type ResponseWithData struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ResponseWithoutData struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func Response(c *gin.Context, code int, message string, payload interface{}) {
	var response interface{}
	if payload != nil {
		response = &ResponseWithData{
			Status:  code,
			Message: message,
			Data:    payload,
		}
	} else {
		response = &ResponseWithoutData{
			Status:  code,
			Message: message,
		}
	}

	c.JSON(code, response)
}
