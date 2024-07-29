package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("content-type", "application/json")
	ctx.JSON(code, gin.H{
		"error":     msg,
		"errorCode": code,
	})
	ctx.Abort()
}

func sendSuccess(ctx *gin.Context, op string, data any) {
	ctx.Header("content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"operation": fmt.Sprintf("operation from handler: %s successful", op),
		"data":      data,
	})
}
