package routers

import (
	"coderminer/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/blog", &controllers.BlogController{})
	beego.Router("/editor", &controllers.EditorController{})
}
