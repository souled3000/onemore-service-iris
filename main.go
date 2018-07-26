package main

import (
	"github.com/kataras/iris"
	myConfig "onemore-service-iris/config"
	"onemore-service-iris/routes"
)

func init() {
	routes.InitRoute()
	myConfig.InitConfig()
}

func main() {
	routes.App.Run(iris.Addr(myConfig.Conf.HttpPort))
}
