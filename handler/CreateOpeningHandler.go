package handler

import (
	"net/http"

	"github.com/CrysLef/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {

	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	err := request.Validate()

	if err != nil {
		logger.Errorf("validation error: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	err = db.Create(&opening).Error

	if err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	SendSuccess(ctx, "create-opening", opening)
}
