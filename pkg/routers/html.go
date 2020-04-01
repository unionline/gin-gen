/*

@Time : 2020/3/20 11:11
@Author : FB
@File : html.go
@Software: GoLand
*/
package routers

import (
	"{{.Appname}}/controllers"

	"github.com/gin-gonic/gin"
)

func routerHTML(r *gin.Engine) {
	// html页面

	tool := r.Group("/")

	tool.GET("/demo/:type", controllers.Html)

}
