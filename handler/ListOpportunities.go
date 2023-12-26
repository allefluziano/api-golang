package handler

import (
	"net/http"

	"github.com/allefluziano/api-golang/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpportunitiesHandler(ctx *gin.Context) {
	opportunities := []schemas.Opportunity{}

	if err := db.Find(&opportunities).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error list opportunities")
		return
	}

	sendSuccess(ctx, "list-opportunities", opportunities)
}
