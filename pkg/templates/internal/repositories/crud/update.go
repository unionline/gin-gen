/*
# @Author : BlackBerry
# @Description : 数据库更新器
# @File : UpdateStyleFabric.go
# @Version : 1.0.0
# @Time : 2019/1/9 14:56
*/

package crud

import (
	"{{.Appname}}/internal/models"
	"github.com/jinzhu/gorm"
)

type Update struct{}

func (Update) UpdateUser(db *gorm.DB, item *models.User) (err error) {
	err = db.Model(item).Update(item).Error
	return
}
