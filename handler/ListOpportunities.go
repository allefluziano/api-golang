package handler

import (
	"net/http"

	"github.com/allefluziano/api-golang/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List Opportunities
// @Description List all Job Opportunities
// @Tags Opportunities
// @Accept json
// @Produce json
// @Success 200 {object} ListOpportunitiesResponse
// @Failure 500 {object} ErrorResponse
// @Router /opportunities [get]
func ListOpportunitiesHandler(ctx *gin.Context) {
	opportunities := []schemas.Opportunity{}

	if err := db.Find(&opportunities).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error list opportunities")
		return
	}

	sendSuccess(ctx, "list-opportunities", opportunities)
}
