package initializers

import (
	"{{.Appname}}/internal/initializers/config"
	"{{.Appname}}/internal/initializers/db"
	"{{.Appname}}/internal/initializers/model"
	"{{.Appname}}/internal/initializers/redis"
)

// Initialize ...
func Initialize() {
	config.Init()
	db.Init()
	redis.Init()
	model.Init()
}
