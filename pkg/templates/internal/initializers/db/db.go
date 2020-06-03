package db

import (
	"fmt"
	"{{.Appname}}/internal/initializers/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"sync"
	
	"github.com/jinzhu/gorm"
	
	// init ...
	_ "github.com/go-sql-driver/mysql"
)

// Db ...
var Db *gorm.DB
var once = &sync.Once{}
var err error

// Init ...
func Init() {
	once.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			config.Setting.MySQL.Username,
			config.Setting.MySQL.Password,
			config.Setting.MySQL.Addr,
			config.Setting.MySQL.DbName,
		)
		
		Db, err = gorm.Open("mysql", dsn)
		if err != nil {
			logrus.Error("db open error:", err)
			return
		}
		
		if err := Db.DB().Ping(); err != nil {
			logrus.Error("db ping error:", err)
			return
		}
		
		Db.DB().SetMaxIdleConns(config.Setting.MySQL.MaxIdleConnections)
		Db.DB().SetMaxOpenConns(config.Setting.MySQL.MaxOpenConnections)
		if config.Setting.GinMode == gin.DebugMode {
			Db.LogMode(true)
		}
	})
}
