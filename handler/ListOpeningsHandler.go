package handler

import (
	"net/http"

	"github.com/CrysLef/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	err := db.Find(&openings).Error
	if err != nil {
		SendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	SendSuccess(ctx, "listing-openings", openings)
}
