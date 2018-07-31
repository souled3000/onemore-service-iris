package utils

import ()

type JSONStruct struct {
	Status string      `json:"status"`
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
	Msg    string      `json:"msg"`
}

func GenRt(status string, code int, rt interface{}, msg string) (r map[string]interface{}) {
	r = make(map[string]interface{})
	r["status"] = status
	r["code"] = code
	r["result"] = rt
	r["msg"] = msg
	return
}
