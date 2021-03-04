package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
)

func Log() gin.HandlerFunc {
	return func(context *gin.Context) {
		path := context.FullPath()
		log.Infof("request path : %v", path)
		context.Next()
	}
}
