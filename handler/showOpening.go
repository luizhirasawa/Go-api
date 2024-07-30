package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizhirasawa/Go-api.git/schemas"
)

// @BasePath /api/v1

// @Summary Show opening
// @Description Show a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [get]
func ShowOpeningHandler(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, "Missing id parameter")
		return
	}
	opening := schemas.Opening{}
	if err := db.Where("id = ?", id).First(&opening).Error; err != nil {
		logger.ErrorF("Error fetching opening: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Error fetching opening from database")
		return
	}
	sendSuccess(c, "showing opening", opening)
}
