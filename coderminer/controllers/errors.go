package controllers

type ErrorController struct {
	BaseController
}

func (this *ErrorController) Error404() {
	this.TplName = "error.html"
}
