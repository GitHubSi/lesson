package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func MiddlePre(ctx *gin.Context) {
	//panic("pre")
	ctx.Next()
}

func MiddleRecover(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("panic occur", r)
		}
	}()

	ctx.Next()
}

var result []int = make([]int, 0, 100)

func main() {
	r := gin.Default()
	//r.Use(MiddlePre, MiddleRecover)

	r.POST("/panic", func(c *gin.Context) {
		panic("this panic")
	})

	result := make(map[int]int)
	r.POST("/crash_panic", func(context *gin.Context) {
		go func() {
			for {
				result[1] = 1
			}
		}()
		_ = result[1]
		context.JSON(http.StatusOK, "")
	})

	r.POST("/timeout", func(context *gin.Context) {
		time.Sleep(time.Second * 4)
		context.JSON(http.StatusOK, "")
	})

	r.Run(":8081")
}
