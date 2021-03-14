package db

import (
	"github.com/prometheus/common/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

//数据库配置
const (
	userName = "mayantao"
	password = "Mayantao110"
	ip       = "120.79.253.180"
	port     = "3306"
	dbName   = "mayantao"
)

var db *gorm.DB

func init() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	db, err := gorm.Open(mysql.Open(path), &gorm.Config{})
	if err != nil {
		log.Info("init db error %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Info("init db error %v", err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func GetDb() *gorm.DB {
	return db
}
