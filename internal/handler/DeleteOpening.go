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
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		SendError(ctx, http.StatusBadRequest, errParamIsRequired("id", typeOf(id)).Error())
		return
	}
	opening := schemas.Opening{}

	err := db.First(&opening, id).Error
	if err != nil {
		SendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	err = db.Delete(&opening).Error
	if err != nil {
		SendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %s", id))
		return
	}

	SendSuccess(ctx, "deleting-opening", opening)
}
