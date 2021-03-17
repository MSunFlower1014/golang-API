package controller

import (
	_ "github.com/MSunFlower1014/golang-API/docs"
	"github.com/MSunFlower1014/golang-API/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// @Summary 查询书籍排行信息
// @Produce  json
// @Param day query string true "指定查询日期，格式yyyyMMdd"
// @Param pageSize query int true "查询条数"
// @Success 200 {string} string  "{"books": [books]}"
// @Router /books [get]
func GetBooksByNow(context *gin.Context) {
	day := context.Query("day")
	pageSizeString := context.Query("pageSize")
	if len(day) == 0 {
		now := time.Now()
		day = now.Format("20060102")
	}
	pageSize := 10
	if len(pageSizeString) > 0 {
		pageSize, _ = strconv.Atoi(pageSizeString)
	}
	books := service.ListBooksByYearMonthDay(day, pageSize)
	context.JSON(http.StatusOK, gin.H{"books": books})
}
