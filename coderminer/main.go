package main

import (
	"coderminer/controllers"
	_ "coderminer/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
