package handler

import (
	"fmt"
	"net/http"

	"github.com/allefluziano/api-golang/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Update Opportunity
// @Description Update a Job Opportunity
// @Tags Opportunities
// @Accept json
// @Produce json
// @Param id query string true "Opportunity Identification"
// @Param opportunity body UpdateOpportunityRequest true "opportunity data to update"
// @Success 200 {object} UpdateOpportunityResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opportunity [put]
func UpdateOpportunityHandler(ctx *gin.Context) {
	request := UpdateOpportunityRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

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

	// Update Opportunity
	if request.Role != "" {
		opportunity.Role = request.Role
	}
	if request.Company != "" {
		opportunity.Company = request.Company
	}
	if request.Location != "" {
		opportunity.Location = request.Location
	}
	if request.Remote != nil {
		opportunity.Remote = *request.Remote
	}
	if request.Link != "" {
		opportunity.Link = request.Link
	}
	if request.Salary > 0 {
		opportunity.Salary = request.Salary
	}

	// Save Opportunity
	if err := db.Save(&opportunity).Error; err != nil {
		logger.Errorf("error updating opportunity: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "erro updating opportunity")
		return
	}

	sendSuccess(ctx, "update-opportunity", opportunity)
}
