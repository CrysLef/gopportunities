package handler

import (
	"fmt"
	"net/http"

	"github.com/CrysLef/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

type HandlerOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

type ListOpeningResponse struct {
	Message string                    `json:"message"`
	Data    []schemas.OpeningResponse `json:"data"`
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"errorCode"`
}

func sendError(ctx *gin.Context, errResponse ErrorResponse) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(errResponse.ErrorCode, errResponse)
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successful", op),
		"data":    data,
	})

}
