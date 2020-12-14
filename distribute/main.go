package main

import (
	"fmt"
	"github.com/think-next/lesson/distribute/router"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		router.GlobalDistribute.HandleChannel(fmt.Sprintf("%d", i), i, router.OperationAdd)
		router.GlobalDistribute.HandleChannel(fmt.Sprintf("%d", i), i, router.OperationRemove)
	}

	time.Sleep(time.Hour)
}
