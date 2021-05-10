package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func main() {
	fmt.Println("vim-go")
	for {
		go func() {
			// 请求第三方服务接口长时间不返回，导致goroutine泄漏
			_, err := http.Get("https://www.baidu.com/")
			if err != nil {
				fmt.Printf("http.Get err: %v\n", err)
			}
		}()

		time.Sleep(time.Second * 1)
		fmt.Println("goroutines: ", runtime.NumGoroutine())
	}
}
