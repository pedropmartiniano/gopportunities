package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pedropmartiniano/gopportunities/schemas"
)

// @BasePath /api/v1

// @Summary Update opening
// @Description Update a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request body"
// @Param id query string true "Opening Id"
// @Success 200 {object} UpdateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/opening [put]
func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		sendError(ctx, 400, err.Error())
		return

	}

	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, 400, errParamIsRequired("id", "queryParam").Error())
		return
	}

	opening := schemas.Opening{}

	err := db.First(opening, id).Error

	if err != nil {
		sendError(ctx, 404, "opening not found")
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
	if request.Remote != nil {
		opening.Remote = request.Remote
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("error updating opening: %v", err.Error())
		sendError(ctx, 500, "error updating opening")
		return
	}

	sendSuccess(ctx, "update-opening", opening)
}
