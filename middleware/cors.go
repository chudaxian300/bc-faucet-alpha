package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		//值可以设为星号,也可以指定具体主机地址,可设置多个地址用逗号隔开,设为指定主机地址第三项才有效
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS")
		//允许请求头修改的类容
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken, X-CSRF-Token, Authorization,Token")
		c.Header("Access-Control-Expose-Headers", "Content-Type,Access-Control-Allow-Origin,Access-Control-Allow-Methods,Access-Control-Allow-Headers,Access-Control-Allow-Credentials")
		//允许使用cookie
		c.Header("Access-Control-Allow-Credentials", "true")

		if  method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}