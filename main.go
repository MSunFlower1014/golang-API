package main

import (
	"fmt"
	"github.com/MSunFlower1014/golang-API/pkg/setting"
	"github.com/MSunFlower1014/golang-API/routers"
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
}
