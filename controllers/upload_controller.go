package controllers

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"io"
	"myblog/models"
	"os"
	"path/filepath"
	"time"
)

type UploadController struct {
	BaseController
}

func (this *UploadController) Post() {
	file, header, err := this.GetFile("upload")
	if err != nil {
		logs.Error(err)
		return
	}
	now := time.Now()
	fileType := "other"
	//判断后缀为图片的文件，如果是图片我们才存入到数据库中
	fileExt := filepath.Ext(header.Filename)
	if fileExt == ".jpg" || fileExt == ".png" || fileExt == ".gif" || fileExt == ".jpeg" {
		fileType = "img"
	}
	//文件夹路径
	fileDir := fmt.Sprintf("static/upload/%s/%d/%d/%d", fileType, now.Year(), now.Month(), now.Day())
	//ModePerm是0777，这样拥有该文件夹路径的执行权限
	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		logs.Error(err)
		return
	}

	//文件路径
	timeStamp := time.Now().Unix()
	fileName := fmt.Sprintf("%d-%s", timeStamp, header.Filename)
	filePathStr := filepath.Join(fileDir, fileName)
	create, err := os.Create(filePathStr)
	if err != nil {
		logs.Error(err)
		return
	}

	//将浏览器客户端上传的文件拷贝到本地路径的文件里面
	_, err = io.Copy(create, file)
	if err != nil {
		logs.Error(err)
		return
	}
	if fileType == "img" {
		album := models.Album{
			Id:         0,
			FilePath:   filePathStr,
			FileName:   header.Filename,
			Status:     0,
			CreateTime: timeStamp,
		}
		_, _ = models.InsertAlbum(album)
	}

	this.Data["json"] = map[string]interface{}{"code": 1, "message": "上传成功"}
	this.ServeJSON()

}
