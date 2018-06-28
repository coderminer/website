package controllers

import (
	"coderminer/models"

	"github.com/globalsign/mgo"
)

var session *mgo.Session = models.GetSession()

type BlogController struct {
	BaseController
}

func (c *BlogController) Get() {
	data, err := models.Blogger.FindBlogById("b261fe1fb938c564c49f782d195709b4")
	if err == nil {
		c.Data["Content"] = data.Content
	}

	c.TplName = "blog.html"
}
