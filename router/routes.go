package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/luizhirasawa/Go-api.git/handler"
)

func initializeRoutes(router *gin.Engine) {
	fmt.Println("Hello, this is a route function.")
	v1 := router.Group("/api/v1")
	{
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.GET("/openings", handler.ListOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
	}
}
