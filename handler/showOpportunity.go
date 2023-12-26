package handler

import (
	"fmt"
	"net/http"

	"github.com/allefluziano/api-golang/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show Opportunity
// @Description Show a Job Opportunity
// @Tags Opportunities
// @Accept json
// @Produce json
// @Param id query string true "Opportunity Identification"
// @Success 200 {object} ShowOpportunityResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opportunity [get]
func ShowOpportunityHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id",
			"queryParameter").Error())
		return
	}

	opportunity := schemas.Opportunity{}

	if err := db.First(&opportunity, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opportunity with id: %s not found", id))
		return
	}

	sendSuccess(ctx, "show-opportunity", opportunity)
}
