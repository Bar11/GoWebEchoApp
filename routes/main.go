package routes

import (
	"GoWebEchoApp/config/setting"
	"GoWebEchoApp/controllers/helloWorld"
	"GoWebEchoApp/controllers/index"
	"GoWebEchoApp/controllers/login"
	"GoWebEchoApp/middlewares"
	"github.com/labstack/echo/v4"
)

func Setup(e *echo.Echo) {
	e.GET(setting.Settings.IndexUrl, index.Index)
	e.GET(setting.Settings.LoginUrl, login.Login)
	e.Use(middlewares.AuthMiddleware) // 以下登录验证
	e.GET("/hello/:name", helloWorld.SayHello)
	e.Logger.Fatal(e.Start("127.0.0.1:8080"))
}
