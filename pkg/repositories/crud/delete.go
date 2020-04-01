/*
# @Author : BlackBerry
# @Description : 数据库删除器
# @File : Delete.go
# @Version : 1.0.0
# @Time : 2019/1/11 14:54
*/

package crud

import (
	"github.com/jinzhu/gorm"
)

type Delete struct{}

// 删除指定ID的记录
func (*Delete) DeleteObjectByID(db *gorm.DB, id uint, model interface{}) (err error) {
	err = db.Where("id = ?", id).Delete(model).Error
	return
}
