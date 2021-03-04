package routers

import (
	"github.com/MSunFlower1014/golang-API/middleware"
	"github.com/MSunFlower1014/golang-API/pkg/setting"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	e := gin.New()
	e.Use(gin.Logger())
	e.Use(gin.Recovery())
	e.Use(middleware.Log())
	gin.SetMode(setting.RunMode)

	e.GET("ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"msg": "pong"})
	})
	return e
}
