package onemore

import (
	"fmt"
	"onemore-service-iris/controllers/base"
	"onemore-service-iris/server"
	"onemore-service-iris/services"
)

func init() {
	server.Srv.Register("version", new(VersionCtl))
}

type VersionCtl struct {
	base.Ctl
}

/**
*
*curl -X POST "http://localhost:8080/api/v2/onemore/version/save" -H "accept: application/json" -H "content-type: application/x-www-form-urlencoded" -d "type=11&app_type=33&html_version=12&html_url=zzzzzz"
**/
func (c *VersionCtl) PostSave() interface{} {
	tp := c.Ctx.FormValue("type")
	at := c.Ctx.FormValue("app_type")
	hv := c.Ctx.FormValue("html_version")
	hu := c.Ctx.FormValue("html_url")
	fmt.Println(tp, at, hv, hu)
	e := services.SaveVersion(tp, at, hv, hu)
	if e != nil {
		fmt.Println("Fail to  save version into redis.", e)
	}
	e = services.SaveVersion2DB(tp, at, hv, hu)
	if e == nil {
		return c.Respond("success", 0, "", "保存成功")
	} else {
		fmt.Println("Fail to  save version into db.", e)
		return c.Respond("error", 1003, "", "保存失败")
	}
}

/**
*curl "localhost:8080/api/v2/onemore/version/getnew?type=11&app_type=22"
**/
func (c *VersionCtl) GetGetnew() interface{} {
	t := c.Ctx.URLParam("type")
	at := c.Ctx.URLParam("app_type")
	fmt.Println(t, at)
	rt, er := services.GetVersion(t, at)
	if er == nil {
		return c.Respond("success", 0, rt, "获取成功")
	} else {
		return c.Respond("error", 1005, nil, "获取失败")
	}
}
