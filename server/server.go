package server

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
	"reflect"
	"strings"
)

var (
	Srv Server
)

func init() {
	Srv.App = iris.New()
	Srv.App.Use(recover.New())
	Srv.App.Use(logger.New())
}

type Server struct {
	App *iris.Application
}

func (this *Server) Register(name string, ctl interface{}) {
	t := reflect.TypeOf(ctl)
	path := t.Elem().PkgPath()
	chips := strings.Split(path, "/")
	chips = chips[2:len(chips)]
	chips = append(chips, name)
	path = strings.Join(chips, "/")
	mvc.New(this.App.Party(path)).Handle(ctl)
}
