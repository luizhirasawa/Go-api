package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizhirasawa/Go-api.git/schemas"
)

// @BasePath /api/v1

// @Summary List openings
// @Description List all job openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningHandler(c *gin.Context) {
	openings := []schemas.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		logger.ErrorF("Error fetching openings: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Error fetching openings from database")
		return
	}
	sendSuccess(c, "listings", openings)
}
