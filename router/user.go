package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitUserRouters() {
	RegistRouter(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		rgPublic.POST("/login", func(c *gin.Context) {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"msg": "login success",
			})
		})
		rgAuthUser := rgAuth.Group("user")
		rgAuthUser.GET("", func(c *gin.Context) {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"data": []map[string]any{
					{"id": 1, "name": "sz"},
					{"id": 2, "name": "ls"},
				},
			})
		})
		rgAuthUser.GET("/:id", func(c *gin.Context) {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"id":   1,
				"name": "sz",
			})
		})
	})
}
