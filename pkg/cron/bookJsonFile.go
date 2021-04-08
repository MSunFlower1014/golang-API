package cron

import (
	"encoding/json"
	"github.com/MSunFlower1014/golang-API/pkg/service"
	"github.com/MSunFlower1014/golang-API/pkg/setting"
	"github.com/prometheus/common/log"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func CreateBookJsonFile() {
	flag := CreateBookDataJson()
	flag = flag && CreateBookLifeJson()
	if flag {
		log.Infof("CreateBookJsonFile execute success")
	} else {
		log.Errorf("CreateBookJsonFile execute fail")
	}
}

func CreateBookDataJson() bool {
	defer func() {
		e := recover()
		if e != nil {
			log.Errorf("CreateBookDataJson recover  err : %v", e)
		}
	}()
	now := time.Now()
	now = now.AddDate(0, -1, 0)
	books := service.ListBooksByYearMonthUnique(now.Year(), int(now.Month()))
	log.Infof("books size is %v", len(*books))
	result := make([]map[string]string, 0)

	for _, book := range *books {
		if book.Name == "" {
			continue
		}
		m := make(map[string]string)
		m["bid"] = book.BID
		m["name"] = book.Name
		m["auth"] = book.Auth
		result = append(result, m)
	}

	bytes, err := json.Marshal(result)
	if err != nil {
		log.Errorf("json marshal error : %v", err)
		return false
	}
	basePath := setting.BookDataJsonPath
	err = ioutil.WriteFile(basePath, bytes, 0666)
	if err != nil {
		log.Errorf("create json file  error : %v", err)
		return false
	}

	return true
}

func CreateBookLifeJson() bool {
	defer func() {
		e := recover()
		if e != nil {
			log.Errorf("CreateBookLifeJson recover  err : %v", e)
		}
	}()
	now := time.Now()
	now = now.AddDate(0, -1, 0)
	books := service.ListBooksByYearMonth(now.Year(), int(now.Month()))
	log.Infof("books size is %v", len(*books))
	result := make([][]string, 1)

	title := []string{"bid", "name", "auth", "rankNum", "time"}
	result[0] = title

	for _, book := range *books {
		temp := make([]string, 0)
		temp = append(temp, strconv.Itoa(book.RankNum))
		temp = append(temp, book.BID)
		temp = append(temp, book.Name)
		temp = append(temp, book.Auth)
		temp = append(temp, book.BYearMonthDay)
		cnt := book.RankCnt
		iCnt := 0
		if strings.Contains(cnt, "万月票") {
			cnt = strings.ReplaceAll(cnt, "万月票", "")
			f, _ := strconv.ParseFloat(cnt, 4)
			iCnt = int(f * 10000)
		} else {
			cnt = strings.ReplaceAll(cnt, "月票", "")
			iCnt, _ = strconv.Atoi(cnt)
		}
		temp = append(temp, strconv.Itoa(iCnt))
		result = append(result, temp)
	}

	bytes, err := json.Marshal(result)
	if err != nil {
		log.Errorf("json marshal err : %v", err)
		return false
	}
	basePath := setting.BookLifeJsonPath
	err = ioutil.WriteFile(basePath, bytes, 0666)
	if err != nil {
		log.Errorf("create json file  error : %v", err)
		return false
	}

	return true
}
