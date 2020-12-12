package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.RedirectTrailingSlash = false

	// 最基础的 GET 请求
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "%s", "hello world")
	})

	r.Run(":8081")
}