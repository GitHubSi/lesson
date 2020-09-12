package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Concurrency()
	}
	wg.Wait()
}

func Concurrency() {
	for {
		HttpDo()
	}
}

func HttpDo() {
	client := &http.Client{}

	req, err := http.NewRequest("POST",
		"http://localhost:8002/initDataByType",
		strings.NewReader("type=lecIndex"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

}
