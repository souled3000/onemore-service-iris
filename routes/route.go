package routes

import (
	"github.com/kataras/iris"
)

var App *iris.Application

func InitRoute() {
	App = iris.New()

	// same as app.Handle("GET", "/ping", [...])
	// Method:   GET
	// Resource: http://localhost:8080/api/v2/onemore/ping
	App.Get("/api/v2/onemore/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello onemore-service-iris!"})
	})
}
