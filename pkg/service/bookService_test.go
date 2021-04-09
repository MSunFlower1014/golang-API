package service

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"runtime/pprof"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestListBooksUnique(t *testing.T) {
	books := ListBooksUnique(0, -2, 0)
	t.Logf("books size is %v", len(*books))
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
		t.Fatalf("json marshal error : %v", err)
	}
	basePath, err := os.Getwd()
	if err != nil {
		t.Fatalf("Getwd error : %v", err)
	}
	path := path.Join(basePath, "book_data.json")
	err = ioutil.WriteFile(path, bytes, 0666)
	if err != nil {
		t.Fatalf("create json file  error : %v", err)
	}
}

func TestListBooksByCreatedTime(t *testing.T) {
	books := ListBooksByGapNowTime(0, -2, 0)
	t.Logf("books size is %v", len(*books))
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
		t.Fatalf("json marshal err : %v", err)
	}
	basePath, err := os.Getwd()
	if err != nil {
		t.Fatalf("Getwd error : %v", err)
	}
	path := path.Join(basePath, "book_life.json")
	err = ioutil.WriteFile(path, bytes, 0666)
	if err != nil {
		t.Fatalf("create json file  error : %v", err)
	}

}

func TestListBooksByYearMonth(t *testing.T) {
	//开启 pprof 性能监控
	cpuf, err := os.Create("cpu_profile")
	if err != nil {
		t.Fatal(err)
	}
	pprof.StartCPUProfile(cpuf)
	defer pprof.StopCPUProfile()

	time.Sleep(time.Second * 3)
	memf, err := os.Create("mem_profile")
	if err != nil {
		t.Fatal("could not create memory profile: ", err)
	}
	if err := pprof.WriteHeapProfile(memf); err != nil {
		t.Fatal("could not write memory profile: ", err)
	}

	memf.Close()

	books := ListBooksByYearMonth(2021, 3)

	for _, v := range *books {
		if v.BYearMonth != "202103" {
			t.Errorf("book info err , BYearMonth is not 202103")
		}
	}
}
