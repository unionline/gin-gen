/*
@Time : 2020/3/12 11:24
@Author : FB
@File : middleware.go
@Software: GoLand
*/
package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Middleware(c *gin.Context) {
	fmt.Println("this is a middleware!")
}
