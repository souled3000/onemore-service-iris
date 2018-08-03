package onemore

import (
	"onemore-service-iris/controllers/base"
	"onemore-service-iris/server"
	"onemore-service-iris/utils"
)

func init() {
	server.Srv.Register("ping", new(PingCtl))
}

type PingCtl struct {
	base.Ctl
}

/*
*curl http://localhost:8080/api/v2/onemore/ping
 */
func (c *PingCtl) Get() interface{} {
	return c.Respond("success", 0, utils.Conf.AppName, "获取成功")
}

//func (c *PingCtl) Get() {
//	c.Ctx.JSON(&base.Response{Status: "success", Code: 0, Result: utils.Conf.AppName, Msg: "获取成功"})
//}
