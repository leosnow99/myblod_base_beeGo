package controllers

import (
	"log"
	"myblog/models"
)

type DeleteArticleController struct {
	BaseController
}

func (this *DeleteArticleController) Get() {
	id, _ := this.GetInt("id")
	_, err := models.DeleteArticle(id)
	if err != nil {
		log.Println(err)
	}
	this.Redirect("/", 302)
}
