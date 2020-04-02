/*
@Time : 2020/3/12 16:50
@Author : FB
@File : model.go
@Software: GoLand
*/
package model

import (
	"{{.Appname}}/initializers/db"
	"{{.Appname}}/models"
)

func Init() {

	if !db.Db.HasTable("users") {
		db.Db.CreateTable(models.User{})
	}
	db.Db.AutoMigrate(models.User{})
}
