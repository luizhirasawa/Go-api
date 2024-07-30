package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizhirasawa/Go-api.git/schemas"
)

func ListOpeningHandler(c *gin.Context) {
	openings := []schemas.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		logger.ErrorF("Error fetching openings: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Error fetching openings from database")
		return
	}
	sendSuccess(c, "listings", openings)
}
