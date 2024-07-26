package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "POST create",
	})
}
