package main

import (
	_ "jms/routers"
	"github.com/astaxie/beego"
)

func main() {
//	beego.SetStaticPath()
	beego.Run()
}

