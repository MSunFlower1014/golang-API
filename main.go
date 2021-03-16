package main

import (
	"fmt"
	cron2 "github.com/MSunFlower1014/golang-API/pkg/cron"
	"github.com/MSunFlower1014/golang-API/pkg/setting"
	"github.com/MSunFlower1014/golang-API/routers"
	"github.com/robfig/cron"
	"log"
	"net/http"
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
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("server start error %v", err)
	}

	c := cron.New()

	_ = c.AddFunc("0 0 0 * * ?", cron2.GetBookRankInfos)
	c.Start()
}
