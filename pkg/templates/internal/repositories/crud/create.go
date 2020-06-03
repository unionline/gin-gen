/*
# @Author : BlackBerry
# @Description : 数据库插入器
# @File : Insert.go
# @Version : 1.0.0
# @Time : 2019/1/8 15:04
*/

package crud

import (
	"{{.Appname}}/internal/models"
	"github.com/jinzhu/gorm"
)

type Create struct{}

func (*Create) CreateUser(db *gorm.DB, item *models.User) (err error) {
	err = db.Create(item).Error
	return
}
