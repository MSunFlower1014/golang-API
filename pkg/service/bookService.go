package service

import (
	"bytes"
	"encoding/json"
	"github.com/MSunFlower1014/golang-API/pkg/dao"
	"github.com/MSunFlower1014/golang-API/pkg/model"
	"github.com/prometheus/common/log"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func SaveBookRank(pageNum, yearmonth string) bool {
	log.Infof("rank book api request pageNum : %v , yearmonth : %v \n", pageNum, yearmonth)
	var buffer bytes.Buffer
	buffer.WriteString("https://m.qidian.com/majax/rank/yuepiaolist?_csrfToken=yOYgIBQMyWxfSQIFmFcanGrSC19FscXSY9qzQXKe&gender=male&pageNum=")
	buffer.WriteString(pageNum)
	buffer.WriteString("&catId=-1&yearmonth=")
	buffer.WriteString(yearmonth)
	url := buffer.String()
	buffer.Reset()
	resp, err := http.Get(url)
	if err != nil {
		log.Infof("fetch: %v\n", err)
		return false
	}

	content, err := ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()

	if err != nil {
		log.Infof("io read all error: %v\n", err)
		return false
	}
	temp, err := zhToUnicode(content)
	if err != nil {
		log.Infof("zhToUnicode error : %v\n", err)
		return false
	}
	var data = make(map[string]interface{})
	if err := json.Unmarshal(temp, &data); err != nil {
		log.Infof("Unmarshal error : %v\n", err)
		return false
	}

	log.Infof("rank book api response code : %v , msg : %s \n", data["code"], data["msg"])

	mainData := data["data"].(map[string]interface{})

	records := mainData["records"].([]interface{})
	now := time.Now()
	yearMonthDay := now.Format("20060102")

	for _, v := range records {
		book := v.(map[string]interface{})
		bookStruct := model.Book{BID: book["bid"].(string), Name: book["bName"].(string),
			Auth: book["bAuth"].(string), BDesc: book["desc"].(string), Cat: book["cat"].(string),
			CatId: int(book["catId"].(float64)), Cnt: book["cnt"].(string), RankCnt: book["rankCnt"].(string),
			RankNum: int(book["rankNum"].(float64)), BYearMonth: yearmonth, BYearMonthDay: yearMonthDay, CreatedAt: now}
		if len(bookStruct.BDesc) > 1000 {
			bookStruct.BDesc = bookStruct.BDesc[0:1000]
		}
		err = dao.InsertBook(&bookStruct)
		if err != nil {
			log.Infof("InsertBook error : %v\n", err)
			return false
		}
	}
	return true
}

func ListBooksByYearMonthDay(yearMontDay string, limit int) *[]model.Book {
	return dao.ListBooksByYearMonthDay(yearMontDay, limit)
}

/*
将json中的unicode转为汉字
*/
func zhToUnicode(raw []byte) ([]byte, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}
