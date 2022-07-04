package router

import (
	"czc-server/app/api"
	"czc-server/app/service"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {

	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(
			service.Middleware.CORS,
		)
		group.ALL("/user", api.UserApi)
	})
}
