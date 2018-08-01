package services

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"onemore-service-iris/misc"
	"onemore-service-iris/models"
	"strconv"
	"strings"
)

func SaveVersion(tp, at, hv, hu string) (err error) {
	r := misc.Redix.Get()
	defer r.Close()
	rt, err := r.Do("hset", "osi_ver", tp+"-"+at, strings.Join([]string{tp, at, hv, hu}, "|"))
	fmt.Println(rt)
	return
}
func SaveVersion2DB(tp, at, hv, hu string) (err error) {
	var o models.Version
	o.AppType, err = strconv.Atoi(at)
	o.Type, err = strconv.Atoi(tp)
	has, err := misc.E.Exist(&o)
	o.HtmlUrl = hu
	o.HtmlVersion = hv
	if has {

		_, err = misc.E.Update(&o, map[string]interface{}{"type": tp, "app_type": at})
	} else {
		_, err = misc.E.InsertOne(&o)
	}
	fmt.Println(err)
	return
}
func GetVersion(tp, at string) (m map[string]string, err error) {
	m = make(map[string]string)
	r := misc.Redix.Get()
	defer r.Close()
	rt, err := redis.String(r.Do("hget", "osi_ver", tp+"-"+at))
	if err != nil {
		o := new(models.Version)
		o.AppType, _ = strconv.Atoi(at)
		o.Type, _ = strconv.Atoi(tp)
		b, err := misc.E.Get(o)
		fmt.Println(o, err, b)
		if err == nil {
			m["type"] = tp
			m["app_type"] = at
			m["html_version"] = o.HtmlVersion
			m["html_url"] = o.HtmlUrl
			return m, nil
		}
		return nil, err
	}
	chips := strings.Split(rt, "|")
	m["type"] = chips[0]
	m["app_type"] = chips[1]
	m["html_version"] = chips[2]
	m["html_url"] = chips[3]
	return
}
