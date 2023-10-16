package handler

import (
	"fmt"
	"net/http"

	"github.com/CrysLef/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Delete opening
// @Description Deleting a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param	id query string true "Opening identification"
// @Success 200 {object} HandlerOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	errResponse := ErrorResponse{}

	if id == "" {
		errResponse.Message = errParamIsRequired("id", typeOf(id)).Error()
		errResponse.ErrorCode = http.StatusBadRequest
		sendError(ctx, errResponse)
		return
	}
	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		errResponse.Message = fmt.Sprintf("opening with id: %s not found", id)
		errResponse.ErrorCode = http.StatusNotFound
		sendError(ctx, errResponse)
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		errResponse.Message = fmt.Sprintf("error deleting opening with id: %s", id)
		errResponse.ErrorCode = http.StatusInternalServerError
		sendError(ctx, errResponse)
		return
	}

	sendSuccess(ctx, "deleting-opening", opening)
}
