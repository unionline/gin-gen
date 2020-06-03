package routers

import (
	"{{.Appname}}/internal/controllers"

	"github.com/gin-gonic/gin"
)

// Initialize ...
func Initialize() *gin.Engine {
	var r = gin.New()

	// 使用Logger中间件
	r.Use(gin.Logger())

	// 使用Recovery中间件
	r.Use(gin.Recovery())

	//加载页面
	r.LoadHTMLGlob("views/template/**/*")


	r.StaticFile("/", "views/index.html")

	// 设置静态文件
	r.Static("static", "views/static")
	r.Static("res", "views/resource")

	// 返回页面，前端js可以计算的路由
	routerHTML(r)

	// 定义路由用来测试使用
	api := r.Group("/api")
	api.GET("users/", controllers.GetUserList)
	api.GET("users/u/:id", controllers.UserInfo)
	api.GET("users/add", controllers.CreateUser)
	api.GET("users/del/", controllers.DeleteUser)
	api.GET("users/update/", controllers.UpdateUser)

	return r
}
