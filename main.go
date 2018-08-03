package main

import (
	"github.com/kataras/iris"
	_ "onemore-service-iris/controllers"
	"onemore-service-iris/server"
	"onemore-service-iris/utils"
)

func main() {
	server.Srv.App.Run(iris.Addr(utils.Conf.HttpPort), iris.WithOptimizations)
}
