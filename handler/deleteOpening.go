package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizhirasawa/Go-api.git/schemas"
)

// @BasePath /api/v1

// @Summary Delete opening
// @Description Delete a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpeningHandler(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, "Missing id parameter")
		return
	}
	opening := schemas.Opening{}
	if err := db.Where("id = ?", id).Delete(&opening).Error; err != nil {
		logger.ErrorF("Error deleting opening: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Error deleting opening from database")
		return
	}
	sendSuccess(c, "deleted", opening)
}
