package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.RedirectTrailingSlash = false

	// 最基础的 GET 请求
	res := r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "%s", "hello")
	})

	// 在上述返回 res 的基础上，继续注册，其实注册的路径还是 /word
	// 而不是 /hello/word
	res.GET("/word", func(c *gin.Context) {
		c.Writer.WriteString("word")
	})
	r.Run(":8081")
}