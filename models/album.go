package models

import (
	"github.com/astaxie/beego/logs"
	"myblog/utils"
)

type Album struct {
	Id         int
	FilePath   string
	FileName   string
	Status     int
	CreateTime int64
}

func InsertAlbum(album Album) (int64, error) {
	sql := "insert into album (filepath, filename, status, createtime) values (?, ?, ?, ?) "
	return utils.ModifyDB(sql, album.FilePath, album.FileName, album.Status, album.CreateTime)
}

func FindAllAlbums() ([]Album, error) {
	rows, err := utils.QueryDB("select id, filepath, filename, status, createtime from album")
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	var albums []Album
	for rows.Next() {
		id := 0
		filepath := ""
		filename := ""
		status := 0
		var createtime int64
		createtime = 0
		_ = rows.Scan(&id, &filepath, &filename, &status, &createtime)
		album := Album{id, filepath, filename, status, createtime}
		albums = append(albums, album)
	}
	return albums, nil
}
