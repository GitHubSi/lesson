package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.RedirectTrailingSlash = true

	// 最基础的 GET 请求
	r.GET("/hello", func(c *gin.Context) {
		//c.String(http.StatusOK, "%s", "hello world")
		type People struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		jsonStr, _ := json.Marshal(People{Name: "fuhui", Age: 1})
		//c.Writer.Write(jsonStr)
		c.JSON(200, jsonStr)
	})

	r.Run(":8081")
}
