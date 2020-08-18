package models

import (
	"fmt"
	"myblog/utils"
)

type User struct {
	Id         int
	Username   string
	Password   string
	Status     int
	Createtime int64
}

//插入
func InsertUser(user User) (int64, error) {
	sql := "insert into users(username, password, status, createtime) values(?, ?, ?, ?)"
	return utils.ModifyDB(sql, user.Username, user.Password, user.Status, user.Createtime)
}

//按条件查询
func QueryUserWithCon(con string) int {
	sql := fmt.Sprintf("select id from users %s ", con)
	row := utils.QueryRowDB(sql)
	id := 0
	_ = row.Scan(&id)
	return id
}

//根据用户名查询
func QueryUserByUsername(username string) int {
	sql := fmt.Sprintf("where username = '%s' ", username)
	return QueryUserWithCon(sql)
}

//根据用户名和密码查询id
func QueryUserWithParam(username, pwd string) int {
	sql := fmt.Sprintf("where username = '%s' and password = '%s' ", username, pwd)
	return QueryUserWithCon(sql)
}
