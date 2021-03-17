package cron

import (
	"github.com/MSunFlower1014/golang-API/pkg/service"
	"github.com/prometheus/common/log"
	"time"
)

const (
	pageNum1 = "1"
	pageNum2 = "2"
	pageNum3 = "3"
)

func GetBookRankInfos() {
	now := time.Now()
	log.Infof("GetBookRankInfos cron start %v", now)
	yearMonthDay := now.Format("200601")
	flag := service.SaveBookRank(pageNum1, yearMonthDay)
	flag = service.SaveBookRank(pageNum2, yearMonthDay) && flag
	//flag = book.SaveBookRank(pageNum3, yearMonthDay) && flag
	log.Infof("GetBookRankInfos cron result %v", flag)

}
