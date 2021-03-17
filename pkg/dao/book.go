package dao

import (
	"github.com/MSunFlower1014/golang-API/pkg/db"
	"github.com/MSunFlower1014/golang-API/pkg/model"
	"github.com/prometheus/common/log"
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
