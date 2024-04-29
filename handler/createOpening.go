package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pedropmartiniano/gopportunities/schemas"
)

// @BasePath /api/v1

// @Summary Create opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request body"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/opening [post]
func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}
	ctx.BindJSON(&request)

	err := request.Validate()
	if err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		sendError(ctx, 400, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	err = db.Create(&opening).Error
	if err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		sendError(ctx, 400, err.Error())
		return
	}

	sendSuccess(ctx, "create-opening", opening)
}
