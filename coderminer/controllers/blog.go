package controllers

import (
	"coderminer/models"
	"fmt"
	"strconv"

	"github.com/globalsign/mgo"
)

var session *mgo.Session = models.GetSession()

type BlogController struct {
	BaseController
}

func (c *BlogController) GetAll() {
	var pageIndex int = 0
	index := c.Ctx.Input.Param(":index")
	if index == "" {
		pageIndex = 0
	} else {
		page, err := strconv.Atoi(index)
		if err != nil {
			pageIndex = 0
		} else {
			pageIndex = page
		}
	}
	data, total, err := models.Blogger.FindAllBlog(pageIndex)
	if err == nil {
		c.Data["Blogs"] = data
		c.Data["Total"] = total
		currentPage := pageIndex + 1
		c.Data["PageNext"] = currentPage
		c.Data["PageBefore"] = pageIndex - 1

		if currentPage != total {
			c.Data["Backward"] = true
		}

		if currentPage != 1 {
			c.Data["Forward"] = true
		}
	}

	c.TplName = "index.html"
}

func (c *BlogController) GetBlogById() {
	id := c.Ctx.Input.Param(":id")
	data, err := models.Blogger.FindBlogById(id)
	if err == nil {
		c.Data["Content"] = data
	} else {
		c.Abort("404")
	}
	fmt.Println("id:", id)

	c.TplName = "detail.html"
}
