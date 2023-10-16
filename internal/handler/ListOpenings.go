package handler

import (
	"net/http"

	"github.com/CrysLef/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary List opening
// @Tags Openings
// @Produce json
// @Success 200 {object} ListOpeningResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}
	errResponse := ErrorResponse{}

	if err := db.Find(&openings).Error; err != nil {
		errResponse.Message = "error listing openings"
		errResponse.ErrorCode = http.StatusInternalServerError
		sendError(ctx, errResponse)
		return
	}

	sendSuccess(ctx, "listing-openings", openings)
}
