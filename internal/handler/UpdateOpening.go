package handler

import (
	"fmt"
	"net/http"

	"github.com/CrysLef/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Update opening
// @Description Updating a new job opening
// @Tags 	Openings
// @Accept 	json
// @Produce json
// @Param 	id query string true "Opening identification"
// @Success 200 {object} UpdateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [put]
func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	err := request.Validate()
	if err != nil {
		logger.Errorf("validation error: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		SendError(ctx, http.StatusBadRequest, errParamIsRequired("id", typeOf(id)).Error())
		return
	}

	opening := schemas.Opening{}

	err = db.First(&opening, id).Error
	if err != nil {
		SendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	err = db.Save(&opening).Error
	if err != nil {
		logger.Errorf("Error updating opening: %v", err.Error())
		SendError(ctx, http.StatusInternalServerError, "error updating opening")
		return
	}

	SendSuccess(ctx, "update-opening", opening)
}
