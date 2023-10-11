package handler

import (
	"fmt"
	"net/http"

	"github.com/CrysLef/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(ctx *gin.Context) {
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

	SendSuccess(ctx, "show-opening", opening)
}
