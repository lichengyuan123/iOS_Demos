package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
)

var RedisTest = redisTestApi{}

type redisTestApi struct{}

func (*redisTestApi) Set(r *ghttp.Request) {
	r.Session.Set("time", gtime.Timestamp())
	r.Response.Write("ok")
}

func (*redisTestApi) Get(r *ghttp.Request) {

	r.Response.Write(r.Session.Map())
}

func (*redisTestApi) Del(r *ghttp.Request) {
	r.Session.Clear()
	r.Response.Write("ok")
}
