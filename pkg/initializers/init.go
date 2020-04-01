package initializers

import (
	"{{.Appname}}/initializers/config"
	"{{.Appname}}/initializers/db"
	"{{.Appname}}/initializers/model"
	"{{.Appname}}/initializers/redis"
)

// Initialize ...
func Initialize() {
	config.Init()
	db.Init()
	redis.Init()
	model.Init()
}
