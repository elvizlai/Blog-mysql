package main

import (
	_ "initial"
	_ "routers"
	_ "functions"
	"github.com/astaxie/beego"
	"controllers"
)

func main() {
	beego.Debug(beego.VERSION)

	beego.Router("/test", &controllers.Test{})
	beego.Router("/upload", &controllers.Upload{})
	beego.Run()
}
