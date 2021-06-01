package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/sync/singleflight"
)

var count int64

func getArticle(id int) (article string, err error) {
	// 假设这里会对数据库进行调用, 模拟不同并发下耗时不同
	atomic.AddInt64(&count, 1)
	time.Sleep(time.Duration(count) * time.Millisecond * 5)
	return fmt.Sprintf("article: %d", id), nil
}

func singleflightGetArticle(sg *singleflight.Group, id int) (string, error) {
	v, err, _ := sg.Do(fmt.Sprintf("%d", id), func() (interface{}, error) {
		return getArticle(id)
	})
	return v.(string), err
}

func main() {
	time.AfterFunc(1*time.Second, func() {
		atomic.AddInt64(&count, -count)
	})

	var (
		wg  sync.WaitGroup
		now = time.Now()
		n   = 10000
		//sg  = &singleflight.Group{}
	)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			// 比较两种方式的不同
			//res, _ := singleflightGetArticle(sg, 1)
			res, _ := getArticle(1)
			if res != "article: 1" {
				panic("err")
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("同时发起 %d 次请求，耗时: %s", n, time.Since(now))
}
