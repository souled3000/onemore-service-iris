package base

import (
	"github.com/kataras/iris"
)

type Ctl struct {
	Ctx iris.Context
}

func (this *Ctl) Respond(status string, code int, rt interface{}, msg string) (r map[string]interface{}) {
	r = make(map[string]interface{})
	r["status"] = status
	r["code"] = code
	r["result"] = rt
	r["msg"] = msg
	return
}

type Response struct {
	Status string      `json:"status"`
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
	Msg    string      `json:"msg"`
}
