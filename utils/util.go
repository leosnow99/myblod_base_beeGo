package utils

import (
	"bytes"
	"crypto/md5"
	"database/sql"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/russross/blackfriday"
	"github.com/sourcegraph/syntaxhighlight"
	"html/template"
)

var db *sql.DB

func InitMysql() {
	fmt.Println("Init Mysql......")
	//注册数据库驱动
	driverName := beego.AppConfig.String("driverName")

	//数据库链接
	user := beego.AppConfig.String("mysqlUser")
	pwd := beego.AppConfig.String("mysqlPassword")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")

	dbConn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		user, pwd, host, port, dbname)

	dbTem, err := sql.Open(driverName, dbConn)
	if err != nil {
		fmt.Println("sql.Open error=", err)
		return
	}
	db = dbTem
	logs.Info("数据库连接成功!")
}

func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		logs.Warning(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		logs.Warning(err)
		return 0, err
	}
	return count, nil

}

func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}

func QueryDB(sql string) (*sql.Rows, error) {
	return db.Query(sql)
}
func MD5(str string) string {
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5str
}

func SwitchMarkdownToHtml(content string) template.HTML {
	markdown := blackfriday.MarkdownCommon([]byte(content))
	//获取到html文档
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(markdown))
	/**
	对document进程查询，选择器和css的语法一样
	第一个参数：i是查询到的第几个元素
	第二个参数：selection就是查询到的元素
	*/
	doc.Find("code").Each(func(i int, selection *goquery.Selection) {
		light, _ := syntaxhighlight.AsHTML([]byte(selection.Text()))
		selection.SetHtml(string(light))
	})
	htmlString, _ := doc.Html()
	return template.HTML(htmlString)
}
