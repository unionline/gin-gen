package main

import (
	"fmt"
	"{{.Appname}}/internal/initializers"
	"{{.Appname}}/internal/initializers/config"
	"{{.Appname}}/internal/routers"
)

func main() {

	// 配置初始化
	initializers.Initialize()

	// 路由初始化
	r := routers.Initialize()

	_ = r.Run(fmt.Sprintf(":%d", config.Setting.Port))
	//r.RunTLS(fmt.Sprintf(":%d", config.Setting.Port), "server.crt", "server.key")

}
