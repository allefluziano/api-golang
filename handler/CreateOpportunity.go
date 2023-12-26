package handler

import (
	"net/http"

	"github.com/allefluziano/api-golang/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Create Opportunity
// @Description Create a new Job Opportunity
// @Tags Opportunities
// @Accept json
// @Produce json
// @Param request body CreateOpportunityRequest true "Request Body"
// @Success 200 {object} CreateOpportunityResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opportunity [post]
func CreateOpportunityHandler(ctx *gin.Context) {
	request := CreateOpportunityRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opportunity := schemas.Opportunity{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opportunity).Error; err != nil {
		logger.Errorf("error creating opportunity: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "create-opportunity", opportunity)
}
