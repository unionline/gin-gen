/*
@Time : 2020/4/1 10:25
@Author : FB
@File : user.go
@Software: GoLand
*/
package repositories

import (
	"{{.Appname}}/initializers/db"
	"{{.Appname}}/models"
	"{{.Appname}}/repositories/crud"
)

var create crud.Create
var read crud.Read
var update crud.Update
var delete crud.Delete

func CreateUser(user *models.User) (err error) {
	tx := db.Db.Begin()
	err = create.CreateUser(tx, user)
	checkTX(tx, "", err)
	return
}

func GetUserList() (users []models.User, err error) {
	tx := db.Db.Begin()
	users, err = read.GetUserList(tx)
	checkTX(tx, "", err)
	return
}

func GetUserInfo(id uint) (user models.User, err error) {
	tx := db.Db.Begin()
	user, err = read.GetUserInfo(tx, id)
	checkTX(tx, "", err)
	return
}

func DeleteUser(id uint) (err error) {
	tx := db.Db.Begin()
	err = delete.DeleteObjectByID(tx, id, models.User{})
	checkTX(tx, "", err)
	return
}

func UpdateUserInfo(item *models.User) (err error) {
	tx := db.Db.Begin()
	err = update.UpdateUser(tx, item)
	checkTX(tx, "", err)
	return
}
