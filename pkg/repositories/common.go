/*
@Time : 2020/4/1 12:24
@Author : FB
@File : common.go
@Software: GoLand
*/
package repositories

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func checkTX(tx *gorm.DB, errmsg string, err error) {
	if err != nil {
		fmt.Println("errmgs=", errmsg)
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}
