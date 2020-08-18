package controllers

import (
	"fmt"
	"myblog/models"
	"time"
)

type ArticleController struct {
	BaseController
}

func (this *ArticleController) Get() {
	this.TplName = "write_article.html"
}

func (this *ArticleController) Post() {
	if !this.IsLogin {
		this.Redirect("/login", 302)
	}
	//获取浏览器传输的数据，通过表单的name属性获取值
	title := this.GetString("title")
	tags := this.GetString("tags")
	short := this.GetString("short")
	content := this.GetString("content")
	fmt.Printf("title:%s,tags:%s\n", title, tags)

	art := models.Article{
		Title:      title,
		Tags:       tags,
		Short:      short,
		Content:    content,
		Author:     this.LoginUser.(string),
		Createtime: time.Now().Unix(),
	}
	_, err := models.AddArticle(art)

	var response map[string]interface{}
	if err == nil {
		//无误
		response = map[string]interface{}{"code": 1, "message": "ok"}
	} else {
		response = map[string]interface{}{"code": 0, "message": "error"}
	}
	this.Data["json"] = response
	this.ServeJSON()

}
