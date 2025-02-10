package main

import (
	models "GoWebEchoApp/models/SQl"
	"GoWebEchoApp/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())  // 使用中间件记录日志
	e.Use(middleware.Recover()) // 使用中间件恢复 panic 到 HTTP 500 响应
	models.InitDB(e)
	routes.Setup(e)                  // 设置路由
	e.Logger.Fatal(e.Start(":8080")) // 启动服务器并监听端口 8080
}
