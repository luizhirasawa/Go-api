package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/luizhirasawa/Go-api.git/handler"

	docs "github.com/luizhirasawa/Go-api.git/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	fmt.Println("Hello, this is a route function.")
	handler.InitHandler()
	BasePath := "/api/v1"
	docs.SwaggerInfo.BasePath = BasePath
	v1 := router.Group(BasePath)
	{
		v1.POST("/opening", handler.PostOpeningHandler)
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.GET("/openings", handler.ListOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
