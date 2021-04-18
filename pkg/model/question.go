package model

import (
	"gorm.io/gorm"
	"time"
)

/*
CREATE TABLE questions
(
	id int primary key AUTO_INCREMENT,
	name varchar(64),
	created_at date,
	answer varchar(2048),
	like_index int,
	show_times int,
	start_level int
)
*/
type Question struct {
	ID         uint `gorm:"primarykey"`
	Name       string
	CreatedAt  time.Time
	Answer     string
	LikeIndex  int
	ShowTimes  int
	StartLevel int
}

func (q *Question) BeforeCreate(tx *gorm.DB) (err error) {
	if len(q.Answer) > 2000 {
		q.Answer = q.Answer[0:1999]
	}
	q.CreatedAt = time.Now()
	return
}
