package routers

import (
	_ "github.com/MSunFlower1014/golang-API/docs"
	"github.com/MSunFlower1014/golang-API/middleware"
	"github.com/MSunFlower1014/golang-API/pkg/controller"
	"github.com/MSunFlower1014/golang-API/pkg/setting"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

func InitRouter() *gin.Engine {
	e := gin.New()
	//gin 设置 pprof
	pprof.Register(e, pprof.DefaultPrefix)

	e.Use(gin.Logger())
	e.Use(gin.Recovery())
	e.Use(middleware.Log())
	gin.SetMode(setting.RunMode)
	e.GET("ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"msg": "pong"})
	})

	e.GET("books", controller.GetBooksByNow)

	e.GET("best", controller.GetBestRankBook)

	e.GET("swagger", Swagger)

	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return e
}

// @Summary swagger测试
// @Produce  json
// @Param name query string true "Name"
// @Success 200 {string} string  "{"msg": "pong"}"
// @Router /swagger [get]
func Swagger(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"msg": "pong"})
}
