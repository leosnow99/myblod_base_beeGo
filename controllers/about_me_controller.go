package controllers

type AboutMeController struct {
	BaseController
}

func (this *AboutMeController) Get() {
	this.Data["wechat"] = "微信：leosnow99"
	this.Data["qq"] = "QQ：905878729"
	this.Data["tel"] = "Tel：null"
	this.TplName = "aboutme.html"
}
