package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"myblog/utils"
	"strconv"
)

type Article struct {
	Id         int
	Title      string
	Tags       string
	Short      string
	Content    string
	Author     string
	Createtime int64
	Status     int //Status=0为正常，1为删除，2为冻结
}

func AddArticle(article Article) (int64, error) {
	i, err := insertArticle(article)
	SetArticleRowsNum()
	return i, err
}

func insertArticle(article Article) (int64, error) {
	sql := "insert into article(title, tags, short, content, author, createtime) values(?,?,?,?,?,?)"
	return utils.ModifyDB(sql, article.Title, article.Tags, article.Short, article.Content, article.Author, article.Createtime)
}

func FindArticleWithPage(page int) ([]Article, error) {
	//从配置文件中获取每页的文章数量
	num, _ := beego.AppConfig.Int("articleListPageNum")
	page--
	fmt.Println("---------->page", page)
	return QueryArticleWithPage(page, num)

}

func QueryArticleWithPage(page, num int) ([]Article, error) {
	sql := fmt.Sprintf("limit %d,%d", page*num, num)
	return QueryArticlesWithCon(sql)
}

func QueryArticlesWithCon(sql string) ([]Article, error) {
	sqlStr := "select id, title, tags, short, content, author, createtime from article " + sql
	rows, err := utils.QueryDB(sqlStr)
	if err != nil {
		fmt.Println("QueryDb error=", err)
		return nil, err
	}

	var artList []Article
	for rows.Next() {
		article := Article{}
		_ = rows.Scan(&article.Id, &article.Title, &article.Tags, &article.Short, &article.Content, &article.Author, &article.Createtime)
		artList = append(artList, article)
	}
	return artList, nil
}

func QueryArticleById(id int) Article {
	sql := "select id,title,tags,short,content,author,createtime from article where id= " + strconv.Itoa(id)
	row := utils.QueryRowDB(sql)
	title := ""
	tags := ""
	short := ""
	content := ""
	author := ""
	var createtime int64
	createtime = 0
	_ = row.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
	art := Article{Id: id, Title: title, Tags: tags, Short: short, Content: content, Author: author, Createtime: createtime}
	return art
}

func UpdateArticle(article Article) (int64, error) {
	sql := "update article set  title=?,tags=?,short=?,content=? where id=?"
	return utils.ModifyDB(sql, article.Title, article.Tags, article.Short, article.Content, article.Id)
}

//----------删除文章---------
func DeleteArticle(artID int) (int64, error) {
	i, err := DeleteArticleById(artID)
	SetArticleRowsNum()
	return i, err
}

func DeleteArticleById(id int) (int64, error) {
	return utils.ModifyDB("delete from article where id = ?", id)
}

func QueryArticleWithParam(param string) []string {
	rows, err := utils.QueryDB(fmt.Sprintf("select %s from article", param))
	if err != nil {
		log.Println(err)
	}
	var paramList []string
	for rows.Next() {
		arg := ""
		_ = rows.Scan(&arg)
		paramList = append(paramList, arg)
	}
	return paramList
}

func QueryArticlesWithTag(tag string) ([]Article, error) {
	sql := " where tags like '%&" + tag + "&%'"
	sql += " or tags like '%&" + tag + "'"
	sql += " or tags like '" + tag + "&%'"
	sql += " or tags like '" + tag + "'"
	fmt.Println(sql)
	return QueryArticlesWithCon(sql)
}
