package main

import (
	_ "study/boot"
	_ "study/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
