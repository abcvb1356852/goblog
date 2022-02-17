package main

import (
	"goblog/bootstrap"
	"goblog/middlewares"
	"goblog/pkg/logger"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()

	err := http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}
