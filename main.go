package main

import (
	"goblog/bootstrap"
	"goblog/config"

	"goblog/middlewares"
	c "goblog/pkg/config"
	"goblog/pkg/logger"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// 初始化配置信息
	config.Initialize()
}

func main() {
	// 初始化 SQL
	bootstrap.SetupDB()

	// 初始化路由绑定
	router := bootstrap.SetupRoute()

	err := http.ListenAndServe(":"+c.GetString("app.port"), middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}
