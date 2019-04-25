package main

import (
	"go-lottery/bootstrap"
	"go-lottery/web/middleware/identity"
	"go-lottery/web/middleware/response"
	"go-lottery/web/routes"
)

func newApp() *bootstrap.Bootstrapper {
	/// 初始化应用

	app := bootstrap.New(
		"Go抽奖系统",
		"Ares",
	)

	app.Bootstrap()
	app.Configure(response.Configure, identity.Configure, routes.Configure)

	return app
}

func main() {
	app := newApp()
	app.Listen(":8080")
}
