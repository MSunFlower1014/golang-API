package dao

import (
	"github.com/MSunFlower1014/golang-API/pkg/db"
	"github.com/MSunFlower1014/golang-API/pkg/model"
)

func InsertQuestion(q *model.Question) error {
	mysqlDb := db.GetDb()
	result := mysqlDb.Create(q)
	return result.Error
}
