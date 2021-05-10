package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	done := make(chan bool)
	arr := []int{1, 2, 3, 4, 5}
	for i, v := range arr {
		b := v
		fmt.Println(i, v, b)
		go func(u int) {
			fmt.Println("Go -> ", u)
			done <- true
		}(v)
	}
	for _ = range arr {
		<-done
	}
}
