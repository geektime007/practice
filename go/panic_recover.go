package main

import (
	"fmt"
	"time"
)

func proc() {
	panic("Ok")
}

func panic_recover() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	proc()
}

func main() {
	fmt.Println("vim-go")
	go func() {
		t := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-t.C:
				go panic_recover()
			}
		}
	}()
	fmt.Println("vim-go-1")
	select {}
}
