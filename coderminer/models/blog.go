package models

import (
	"fmt"

	"github.com/globalsign/mgo/bson"

	"github.com/globalsign/mgo"
)

const (
	database = "Blog"
	dbName   = "BlogModel"
	limit    = 10
)

type Blog struct {
	Id       string `bson:"_id"`
	Title    string `bson:"title"`
	Author   string `bson:"author"`
	Des      string `bson:"des"`
	Img      string `bson:"img"`
	Original string `bson:"original"`
	Content  string `bson:"content"`
	DateStr  string `bson:"dateStr"`
	Date     int64  `bson:"date"`
	Tags     []Tag  `bson:"tags"`
}

type Tag struct {
	Alias string `bson:"alias"`
	Name  string `bson:"name"`
}

var session *mgo.Session

func NewBlog() *Blog {
	session = GetSession()

	return &Blog{}
}

func (b *Blog) PostBlog(blog *Blog) error {
	sc := session.Copy()
	defer sc.Close()
	c := sc.DB(database).C(dbName)

	err := c.Insert(blog)

	if err != nil {
		fmt.Println("insert model error: ", err)
	}

	return err
}

func (b *Blog) FindBlogById(id string) (Blog, error) {
	sc := session.Copy()
	defer sc.Close()
	c := sc.DB(database).C(dbName)

	var blog Blog
	err := c.Find(bson.M{"_id": id}).One(&blog)
	if err != nil {
		fmt.Println("FindById error", err)
	}
	return blog, err
}

func (b *Blog) FindBlogByTag(tag string, pageindex int) ([]Blog, error) {
	sc := session.Copy()
	defer sc.Close()
	c := sc.DB(database).C(dbName)

	var blogs []Blog
	err := c.Find(bson.M{"tags.name": tag}).Skip(pageindex).Limit(limit).All(&blogs)
	return blogs, err
}

func (b *Blog) FindAllBlog(pageindex int) ([]Blog, int, error) {
	sc := session.Copy()
	defer sc.Close()
	var blogs []Blog

	c := sc.DB(database).C(dbName)
	total, _ := c.Count()
	totalPages := total/limit + 1
	err := c.Find(nil).Skip(pageindex * limit).Limit(limit).Sort("-date").All(&blogs)

	return blogs, totalPages, err
}
