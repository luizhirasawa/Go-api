package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizhirasawa/Go-api.git/schemas"
)

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
