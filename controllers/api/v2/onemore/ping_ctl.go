package onemore

import (
	"onemore-service-iris/config"
	"onemore-service-iris/server"
	"onemore-service-iris/utils"
)

func init() {
	server.Srv.Register("ping", new(PingCtl))
}

type PingCtl struct {
}

func (c *PingCtl) Get() interface{} {
	return utils.GenRt("success", 0, config.Conf.AppName, "获取成功")
}
