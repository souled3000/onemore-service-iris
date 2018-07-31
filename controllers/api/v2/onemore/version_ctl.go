package onemore

import (
	"onemore-service-iris/server"
	//	"onemore-service-iris/utils"
)

func init() {
	server.Srv.Register("ping", new(VersionCtl))
}

type VersionCtl struct {
}
