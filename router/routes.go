package router

import (
	"github.com/allefluziano/api-golang/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize Handler
	handler.InitializeHandler()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/opportunity", handler.ShowOpportunityHandler)

		v1.POST("/opportunity", handler.CreateOpportunityHandler)

		v1.DELETE("/opportunity", handler.DeleteOpportunityHandler)

		v1.PUT("/opportunity", handler.UpdateOpportunityHandler)

		v1.GET("/opportunities", handler.ListOpportunitiesHandler)
	}
}
