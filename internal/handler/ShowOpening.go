package handler

import (
	"fmt"
	"net/http"

	"github.com/CrysLef/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Show opening
// @Description Showing a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param	id query string true "Opening identification"
// @Success 200 {object} HandlerOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [get]
func ShowOpeningHandler(ctx *gin.Context) {
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

	sendSuccess(ctx, "show-opening", opening)
}
