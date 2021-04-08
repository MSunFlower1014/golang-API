package main

import (
	"fmt"
	cron2 "github.com/MSunFlower1014/golang-API/pkg/cron"
	"github.com/MSunFlower1014/golang-API/pkg/setting"
	"github.com/MSunFlower1014/golang-API/routers"
	"github.com/prometheus/common/log"
	"github.com/robfig/cron"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:              fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:           router,
		ReadHeaderTimeout: setting.ReadTimeout,
		WriteTimeout:      setting.WriteTimeout,
		MaxHeaderBytes:    1 << 20,
	}

	CronInit()
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("server start error %v", err)
	}

}

//初始化定时任务
func CronInit() {
	c := cron.New()

	err := c.AddFunc("0 1 0 * * *", cron2.GetBookRankInfos)
	if err != nil {
		log.Errorf("add cron GetBookRankInfos error %v", err)
	}

	err = c.AddFunc("0 1 0 1 * *", cron2.CreateBookJsonFile)
	if err != nil {
		log.Errorf("add cron CreateBookJsonFile error %v", err)
	}
	c.Start()
}
