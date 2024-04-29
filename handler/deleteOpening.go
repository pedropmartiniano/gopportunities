package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pedropmartiniano/gopportunities/schemas"
)

// @BasePath /api/v1

// @Summary Delete opening
// @Description Delete a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening Id"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/opening [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
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

	err = db.Delete(&opening, id).Error

	if err != nil {
		sendError(ctx, 500, fmt.Sprintf("error deleting opening with id %s", id))
		return
	}

	sendSuccess(ctx, "delete-opening", opening)
}
