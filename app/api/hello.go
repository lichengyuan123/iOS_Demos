package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Hello = helloApi{}

type helloApi struct{}

// Index is a demonstration route handler for output "Hello World!".
func (*helloApi) Index(r *ghttp.Request) {
	// for i := 1; i < 1000000; i++ {
	// 	g.Model("test1").Data(g.Map{"id":i}).Insert()
	// }
	// timeNow := gtime.Millisecond()
	for i := 1; i < 1000000; i++ {
		g.Model("test3").Data(g.Map{"id": i}).Insert()
	}

	// println(timeNow)

	r.Response.Writeln("Hello World!")
}
