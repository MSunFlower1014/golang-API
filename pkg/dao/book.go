package dao

import (
	"github.com/MSunFlower1014/golang-API/pkg/db"
	"github.com/MSunFlower1014/golang-API/pkg/model"
	"github.com/prometheus/common/log"
	"time"
)

func InsertBook(book *model.Book) error {
	mysqlDB := db.GetDb()
	result := mysqlDB.Create(book)
	return result.Error
}

func ListBooksByYearMonthDay(yearMontDay string, limit int) *[]model.Book {
	var books []model.Book
	mysqlDB := db.GetDb()
	result := mysqlDB.Where("b_year_month_day = ?", yearMontDay).Limit(limit).Find(&books)
	if result.Error != nil {
		log.Errorf("ListBooksByYearMonthDay error : %v", result.Error)
	}
	return &books
}

func ListFirstRankBookByLimitDays(rankNum, limit int) *[]model.Book {
	var books []model.Book
	mysqlDB := db.GetDb()
	result := mysqlDB.Where("rank_num = ?", rankNum).Order("id desc").Limit(limit).Find(&books)
	if result.Error != nil {
		log.Errorf("ListFirstRankBookByLimitDays error : %v", result.Error)
	}
	return &books
}

func DeleteBookByName(name string) (int64, error) {
	mysqlDB := db.GetDb()
	result := mysqlDB.Where("name = ?", "test").Delete(model.Book{})
	if result.Error != nil {
		log.Errorf("DeleteBookByName error : %v", result.Error)
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func ListBooksByCreatedTime(year, month, day int) *[]model.Book {
	var books []model.Book
	mysqlDB := db.GetDb()
	now := time.Now()
	now = now.AddDate(year, month, day)
	result := mysqlDB.Where("created_at > ?", now).Order("b_year_month_day,rank_num").Find(&books)
	if result.Error != nil {
		log.Errorf("ListFirstRankBookByLimitDays error : %v", result.Error)
	}
	return &books
}
