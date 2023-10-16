package handler

import (
	"net/http"

	"github.com/CrysLef/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Create opening
// @Description Creating a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body HandlerOpeningRequest true "Request body"
// @Success 200 {object} HandlerOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func CreateOpeningHandler(ctx *gin.Context) {
	opening := schemas.Opening{}
	request := HandlerOpeningRequest{}
	errResponse := ErrorResponse{}

	ctx.BindJSON(&request)

	if err := request.ValidateCreate(); err != nil {
		errResponse.Message = err.Error()
		errResponse.ErrorCode = http.StatusBadRequest
		logger.Errorf("validation error: %v", errResponse.Message)
		sendError(ctx, errResponse)
		return
	}

	opening = schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		errResponse.Message = err.Error()
		errResponse.ErrorCode = http.StatusInternalServerError
		logger.Errorf("error creating opening: %v", errResponse.Message)
		sendError(ctx, errResponse)
		return
	}
	sendSuccess(ctx, "create-opening", opening)
}
