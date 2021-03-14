package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/MSunFlower1014/golang-API/pkg/dao"
	"github.com/MSunFlower1014/golang-API/pkg/model"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func BooksInsert(pageNum, month string) {

	var buffer bytes.Buffer
	buffer.WriteString("https://m.qidian.com/majax/rank/yuepiaolist?_csrfToken=yOYgIBQMyWxfSQIFmFcanGrSC19FscXSY9qzQXKe&gender=male&pageNum=")
	buffer.WriteString(pageNum)
	buffer.WriteString("&catId=-1&yearmonth=")
	buffer.WriteString(month)
	url := buffer.String()
	buffer.Reset()
	resp, err := http.Get(url)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		return
	}

	content, err := ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		return
	}
	//fmt.Printf("%s", content)
	temp, err := zhToUnicode(content)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		return
	}
	//fmt.Printf("%s", temp)
	var data = make(map[string]interface{})
	if err := json.Unmarshal(temp, &data); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		return
	}

	fmt.Printf("code : %v , msg : %s \n", data["code"], data["msg"])

	mainData := data["data"].(map[string]interface{})

	records := mainData["records"].([]interface{})
	now := time.Now()
	yearMonthDay := time.Now().Format("20060102")

	for _, v := range records {
		book := v.(map[string]interface{})
		bookStruct := model.Book{&gorm.Model{}, book["bid"].(string), book["bName"].(string),
			book["bAuth"].(string), book["desc"].(string), book["cat"].(string),
			int(book["catId"].(float64)), book["rankCnt"].(string),
			int(book["rankNum"].(float64)), month, yearMonthDay}
		if len(bookStruct.Desc) > 1000 {
			bookStruct.Desc = bookStruct.Desc[0:1000]
		}
		err := dao.InsertBook(&bookStruct)
		if err != nil {

		}
	}

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
