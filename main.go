package main

import (
	"github.com/kataras/iris"
	"onemore-service-iris/config"
	_ "onemore-service-iris/controllers"
	_ "onemore-service-iris/misc"
	"onemore-service-iris/server"
)

func main() {
	server.Srv.App.Get("/api/v2/onemore/ping2", f)
	server.Srv.App.Run(iris.Addr(config.Conf.HttpPort), iris.WithOptimizations)
}

func f(ctx iris.Context) {
	ctx.JSON(iris.Map{"status": "succes", "code": 0, "result": config.Conf.AppName, "msg": "获取成功"})
}
