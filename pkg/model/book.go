package model

import (
	"gorm.io/gorm"
	"time"
)

/*
CREATE TABLE books
(
	id int primary key AUTO_INCREMENT,
	created_at date,
	updated_at date,
	b_id varchar(64),
	name varchar(64),
	auth varchar(64),
 	b_desc varchar(1000),
	cat varchar(64),
	cat_id int ,
	cnt varchar(64),
	rank_cnt varchar(64),
	rank_num int,
	b_year_month varchar(64),
	b_year_month_day varchar(64)
)
*/
type Book struct {
	ID            uint `gorm:"primarykey"`
	BID           string
	Name          string
	Auth          string
	BDesc         string
	Cat           string
	CatId         int
	Cnt           string
	RankCnt       string
	RankNum       int
	BYearMonth    string
	BYearMonthDay string
	CreatedAt     time.Time
}

func (book *Book) BeforeCreate(tx *gorm.DB) (err error) {
	if len(book.BDesc) > 1000 {
		book.BDesc = book.BDesc[0:999]
	}
	return
}
