package router

import (
	"fmt"

	"github.com/luizhirasawa/Go-api.git/docs"

	"github.com/gin-gonic/gin"
	"github.com/luizhirasawa/Go-api.git/handler"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	fmt.Println("Hello, this is a route function.")
	handler.InitHandler()
	BasePath := "/api/v1"
	docs.SwaggerInfo.BasePath = BasePath
	v1 := router.Group(BasePath)
	path := "/opening"
	{
		v1.POST(path, handler.PostOpeningHandler)
		v1.DELETE(path, handler.DeleteOpeningHandler)
		v1.GET(path, handler.ShowOpeningHandler)
		v1.GET("/openings", handler.ListOpeningHandler)
		v1.PUT(path, handler.UpdateOpeningHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
