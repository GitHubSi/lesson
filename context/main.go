package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)
	go func() {
		for {
			select {
			default:
				fmt.Println("run")
				time.Sleep(time.Second)
			case <-ctx.Done():
				fmt.Println(ctx.Err())
				return
			}
		}
	}()

	time.Sleep(time.Second * 5)
	cancel()
}
