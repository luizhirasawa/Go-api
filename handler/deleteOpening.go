package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizhirasawa/Go-api.git/schemas"
)

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
