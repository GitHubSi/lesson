package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jpillora/overseer"
	"net/http"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {
	fmt.Println("middle write log")
	return func(c *gin.Context) {
		fmt.Println("again")
	}
}

var server = &http.Server{}
var addresses = make([]string, 0)

// 平滑重启
func GraceStart(addr string, gin *gin.Engine) {
	addresses = append(addresses, addr)
	server = &http.Server{
		Addr:    addr,
		Handler: gin,
	}
	overseer.Run(overseer.Config{
		Program:   prog,
		Addresses: addresses,
	})
}

func prog(state overseer.State) {
	server.Serve(state.Listener)
}

// overseer 如果只启动的是HTTP的服务进程，那么在Main下的常驻协程会执行两遍
// overseer 执行平滑重启的话，Main下的常驻协程不会被更新，除非程序重新重启，否则配置永远不会被更新；
func main() {
	router := gin.New()
	router.Use(LoggerMiddleware())

	// 服务的常驻协成程
	go func() {

		// 执行两次
		fmt.Println("start go routine")
		time.Sleep(time.Minute * 2)
		fmt.Println("end upgrade routine")
	}()

	router.GET("/get", func(c *gin.Context) {
		fmt.Println("get response")
	})

	router.POST("/post", func(c *gin.Context) {
		fmt.Println("post response ")
	})

	GraceStart(":1313", router)
}
