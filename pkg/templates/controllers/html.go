/*
@Time : 2020/3/20 11:08
@Author : FB
@File : html.go
@Software: GoLand
*/
package controllers

import (
	"fmt"
	"{{.Appname}}/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Html(ctx *gin.Context) {

	enc_type := ctx.Param("type")
	fmt.Println("type=", enc_type)

	uri := ctx.Request.RequestURI

	var name string

	name = uri
	if strings.LastIndex(uri, ".html") == -1 {
		name = uri + ".html"
	}

	fmt.Println("uri after=", uri)
	err := utils.ValidTemplateName(name)
	if err != nil {

		ctx.HTML(http.StatusOK, "redirect.html", gin.H{
			"title": "跳转首页",
		})
		return
	}

	ctx.HTML(http.StatusOK, name, gin.H{
		"title": "sample base64",
	})
}
