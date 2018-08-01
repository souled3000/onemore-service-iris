package onemore

import (
	"onemore-service-iris/config"
	"onemore-service-iris/misc"
	"onemore-service-iris/server"
)

func init() {
	server.Srv.Register("ping", new(PingCtl))
}

type PingCtl struct {
}

func (c *PingCtl) Get() interface{} {
	return misc.GenRt("success", 0, config.Conf.AppName, "获取成功")
}
