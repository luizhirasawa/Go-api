package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoute(router *gin.Engine) {
	fmt.Println("Hello, this is a route function.")
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "opening",
			})
		})
		v1.POST("/opening", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "opening",
			})
		})
		v1.PUT("/opening", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "opening",
			})
		})
		v1.DELETE("/opening", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "opening",
			})
		})
		v1.GET("/openings", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "opening",
			})
		})
	}
}
