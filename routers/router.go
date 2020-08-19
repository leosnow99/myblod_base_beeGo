package routers

import (
	"github.com/astaxie/beego"
	"myblog/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/exit", &controllers.LoginController{})
	beego.Router("/article/add", &controllers.ArticleController{})
	//显示文章内容
	beego.Router("/article/:id", &controllers.ShowArticleController{})
	beego.Router("/article/update", &controllers.UpdateArticleController{})
	beego.Router("/article/delete", &controllers.DeleteArticleController{})
	beego.Router("/tags", &controllers.TagsController{})
	//相册
	beego.Router("/album", &controllers.AlbumController{})

	//文件上传
	beego.Router("/upload", &controllers.UploadController{})
	beego.Router("/aboutme", &controllers.AboutMeController{})
}
