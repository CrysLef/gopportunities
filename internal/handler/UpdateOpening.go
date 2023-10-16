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
// @Success 200 {object} HandlerOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [put]
func UpdateOpeningHandler(ctx *gin.Context) {
	opening := schemas.Opening{}
	request := HandlerOpeningRequest{}
	errResponse := ErrorResponse{}

	ctx.BindJSON(&request)

	if err := request.ValidateUpdate(); err != nil {
		errResponse.Message = err.Error()
		errResponse.ErrorCode = http.StatusBadRequest

		logger.Errorf("validation error: %v", errResponse.Message)
		sendError(ctx, errResponse)
		return
	}

	id := ctx.Query("id")
	if id == "" {
		errResponse.Message = errParamIsRequired("id", typeOf(id)).Error()
		errResponse.ErrorCode = http.StatusBadRequest

		sendError(ctx, errResponse)
		return
	}

	if err := db.First(&opening, id).Error; err != nil {
		errResponse.Message = fmt.Sprintf("opening with id: %s not found", id)
		errResponse.ErrorCode = http.StatusNotFound
		sendError(ctx, errResponse)
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

	if request.Remote || !request.Remote {
		opening.Remote = request.Remote
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	if err := db.Save(&opening).Error; err != nil {
		errResponse.Message = "error updating opening"
		errResponse.ErrorCode = http.StatusInternalServerError
		logger.Errorf("Error updating opening: %v", err.Error())
		sendError(ctx, errResponse)
		return
	}

	sendSuccess(ctx, "update-opening", opening)
}
