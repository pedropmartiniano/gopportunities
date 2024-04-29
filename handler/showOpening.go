package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pedropmartiniano/gopportunities/schemas"
)

// @BasePath /api/v1

// @Summary Show opening
// @Description Show a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening Id"
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/opening [get]
func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, 400, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	err := db.First(&opening, id).Error

	if err != nil {
		sendError(ctx, 404, fmt.Sprintf("opening with id %s not found", id))
		return
	}

	sendSuccess(ctx, "show-opening", opening)
}
