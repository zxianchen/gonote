package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type IFnRegistRouter = func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup)

var (
	gfnRouters []IFnRegistRouter
)

func RegistRouter(fn IFnRegistRouter) {
	if fn == nil {
		return
	}
	gfnRouters = append(gfnRouters, fn)
}
func InitRouter() {
	r := gin.Default()

	rgPublic := r.Group("/api/v1/public")
	rgAuth := r.Group("/api/v1")
	InitBasePlatformRouters()
	for _, fn := range gfnRouters {
		fn(rgPublic, rgAuth)
	}
	port := viper.GetString("server.port")
	if port == "" {
		port = "8090"
	}
	err := r.Run(":" + port)
	if err != nil {
		panic(fmt.Sprint("Start server error: %s", err.Error()))
	}
}
func InitBasePlatformRouters() {
	InitUserRouters()
}
