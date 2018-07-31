package main

import (
	"github.com/kataras/iris"
	"onemore-service-iris/routes"
	"onemore-service-iris/utils"
)

func init() {
	routes.InitRoute()
	utils.InitConfig()
}

func main() {
	routes.App.Run(iris.Addr(utils.Conf.HttpPort))
}
