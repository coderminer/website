package routers

import (
	"coderminer/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/blog/?:index(^[0-9]*$)", &controllers.BlogController{}, "get:GetAll")
	beego.Router("/blog/:id", &controllers.BlogController{}, "get:GetBlogById")
	beego.Router("/editor", &controllers.EditorController{})
}
