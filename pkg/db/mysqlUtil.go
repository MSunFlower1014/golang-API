package db

import (
	"github.com/prometheus/common/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strings"
	"time"
)

//数据库配置
const (
	userName = "root"
	password = "toor"
	ip       = "120.79.253.180"
	port     = "3307"
	dbName   = "mayantao"
)

var db *gorm.DB

/*
error :
sql: Scan error on column index 2, name “created_at“: unsupported Scan
需要在链接地址后增加parseTime=true
https://blog.csdn.net/galoiszhou/article/details/114257575
*/
func init() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8&parseTime=true"}, "")
	var err error
	db, err = gorm.Open(mysql.Open(path), &gorm.Config{
		//设置日志级别
		Logger: logger.Default.LogMode(logger.Info),
	})
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
