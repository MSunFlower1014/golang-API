package model

import (
	"gorm.io/gorm"
)

/*
CREATE TABLE BOOKS
(
	ID varchar(32)
)
*/
type Book struct {
	gorm.Model
	BID          string
	Name         string
	Auth         string
	Desc         string
	Cat          string
	CatId        int
	Cnt          string
	RankCnt      string
	RankNum      int
	YearMonth    string
	YearMonthDay string
}
