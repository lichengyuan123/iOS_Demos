package router

import (
	"study/app/api"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gsession"
)

func init() {
	s := g.Server()
	s.SetConfigWithMap(g.Map{
		"SessionMaxAge":  time.Minute,
		"SessionStorage": gsession.NewStorageRedis(g.Redis()),
	})

	s.Group("/", func(group *ghttp.RouterGroup) {

		group.ALL("/redisTest", api.RedisTest)

		group.ALL("/hello", api.Hello)

		
	})

	s.SetPort(8199)
}
