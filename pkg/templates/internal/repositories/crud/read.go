/*
@Time : 2020/4/1 10:33
@Author : FB
@File : read.go
@Software: GoLand
*/
package crud

import (
	"{{.Appname}}/internal/models"
	"github.com/jinzhu/gorm"
)

type Read struct{}

func (Read) GetUserInfo(db *gorm.DB, id uint) (out models.User, err error) {
	//err = db.Where("id=?",id).First(&out).Error
	err = db.First(&out, id).Error
	return
}

func (Read) GetUserList(db *gorm.DB) (out []models.User, err error) {
	//err = db.Where("id=?",id).First(&out).Error
	err = db.Find(&out).Error
	return
}
