package main

import (
	"blog/app/http/middlewares"
	"blog/bootstrap"
	"blog/config"
	c "blog/pkg/config"
	"net/http"
)

func init() {
	// 初始化配置信息
	config.Initialize()
}

func main() {

	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()

	http.ListenAndServe(":"+c.GetString("app.port"), middlewares.RemoveTrailingSlash(router))
}
