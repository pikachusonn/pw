package httpserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type SuccessAPIResponse struct {
	Status int `json:"status"`
	Data   any `json:"data"`
}

type Failure struct {
	Error APIError `json:"error"`
}

func OK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, SuccessAPIResponse{
		Status: http.StatusOK,
		Data:   data,
	})
}

func Created(c *gin.Context, data any) {
	c.JSON(http.StatusCreated, SuccessAPIResponse{
		Status: http.StatusCreated,
		Data:   data,
	})
}

func WriteError(c *gin.Context, code string, status int, message string) {
	c.AbortWithStatusJSON(status, Failure{
		Error: APIError{
			Code:    code,
			Message: message,
		},
	})
}
