package controllers

import (
	"coderminer/helper"
	"coderminer/models"
	"encoding/json"
	"time"

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
	Title  string `json:"title"`
	Des    string `json:"des"`
	Img    string `json:"Img"`
}

func (c *EditorController) Post() {

	var data Object
	req := string(c.Ctx.Input.RequestBody)
	json.Unmarshal([]byte(req), &data)

	blogger := models.NewBlog()
	blogger.Id = helper.GetMd5(bson.NewObjectId().Hex())

	blogger.Content = data.Html
	blogger.Des = data.Html[0:200]
	blogger.Title = data.Title
	blogger.DateStr = time.Now().Format("2006年01月02日")
	blogger.Author = "Kevin"
	blogger.Des = data.Des
	blogger.Img = data.Img
	blogger.Date = time.Now().Unix()
	blogger.PostBlog(blogger)

	//fmt.Println("data",data.Html)

	c.ServeJSON()
}
