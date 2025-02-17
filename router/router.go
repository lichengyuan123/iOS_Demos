package router

import (
	"study/app/api"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()

	s.Group("/", func(group *ghttp.RouterGroup) {

		group.ALL("/hello", api.Hello)

	})

	s.SetPort(8199)
}
