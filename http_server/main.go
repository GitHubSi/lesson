package main

import (
    "encoding/json"
    "github.com/gin-gonic/gin"
    "github.com/think-next/lesson/trace/trace"
)

func GinMiddleWare() gin.HandlerFunc {
    return func(context *gin.Context) {
        defer trace.Trace()()

        // context.Next()
        // context.AbortWithStatusJSON(http.StatusForbidden, gin.H{"code": 1})
    }
}

func GinAuth() gin.HandlerFunc {
    return func(context *gin.Context) {
        defer trace.Trace()()

        // context.Next()
        // context.AbortWithStatusJSON(http.StatusForbidden, gin.H{"code": 1})
    }
}

func main() {
    r := gin.Default()
    r.RedirectTrailingSlash = true

    r.Use(GinMiddleWare(), GinAuth())

    // 最基础的 GET 请求
    r.GET("/hello", func(c *gin.Context) {
        defer trace.Trace()()

        // c.String(http.StatusOK, "%s", "hello world")
        type People struct {
            Name string `json:"name"`
            Age  int    `json:"age"`
        }
        jsonStr, _ := json.Marshal(People{Name: "fuhui", Age: 1})
        // fmt.Println("processing")
        // c.Writer.Write(jsonStr)
        c.JSON(200, jsonStr)
    })

    r.Run(":8081")
}
