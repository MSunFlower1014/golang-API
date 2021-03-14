package dao

import (
	"github.com/MSunFlower1014/golang-API/pkg/db"
	"github.com/MSunFlower1014/golang-API/pkg/model"
)

func InsertBook(book *model.Book) error {
	mysqlDB := db.GetDb()
	result := mysqlDB.Create(book)
	return result.Error
}
