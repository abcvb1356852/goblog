package model

import (
	"goblog/pkg/logger"

	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	// GORM 的 MySQL 数据库驱动导入
	"gorm.io/driver/mysql"
)

// DB gorm.DB 对象
var DB *gorm.DB

// ConnectDB 初始化模型
func ConnectDB() *gorm.DB {

	var err error

	config := mysql.New(mysql.Config{
		DSN: "root:sj19920123@tcp(116.63.156.28:33061)/goblog?charset=utf8&parseTime=True&loc=Local",
	})

	/* 	// 准备数据库连接池
	   	DB, err = gorm.Open(config, &gorm.Config{}) */
	// 准备数据库连接池
	DB, err = gorm.Open(config, &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Warn),
	})

	logger.LogError(err)

	return DB
}