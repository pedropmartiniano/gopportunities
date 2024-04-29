package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pedropmartiniano/gopportunities/schemas"
)

// @BasePath /api/v1

// @Summary List opening
// @Description List a job openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/openings [get]
func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	err := db.Find(&openings).Error

	if err != nil {
		sendError(ctx, 500, "error listing openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
