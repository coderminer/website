package controllers

import (
	"coderminer/helper"
	"coderminer/models"
	"encoding/json"

	"github.com/globalsign/mgo/bson"
)

type EditorController struct {
	BaseController
}

func (c *EditorController) Get() {
	c.TplName = "editor.html"
}

type Object struct {
	Origin string `json:"origin"`
	Html   string `json:"html"`
}

func (c *EditorController) Post() {

	var data Object
	req := string(c.Ctx.Input.RequestBody)
	json.Unmarshal([]byte(req), &data)

	blogger := models.NewBlog()
	blogger.Id = helper.GetMd5(bson.NewObjectId().Hex())

	blogger.Original = data.Origin
	blogger.Content = data.Html
	blogger.PostBlog(blogger)

	//fmt.Println("data",data.Html)

	c.ServeJSON()
}
