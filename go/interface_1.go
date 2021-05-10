package main

import "fmt"

// 空接口
type A interface{}

func main() {
	fmt.Println("vim-go")
	var a A
	var str = "你好Golang"
	a = str // 让字符串实现A接口
	fmt.Printf("值: %v, 类型: %T\n", a, a)

	var num = 20
	a = num // 让int类型实现A接口
	fmt.Printf("值: %v, 类型: %T\n", a, a)

	var flag = true
	a = flag
	fmt.Printf("值: %v, 类型: %T", a, a)
}
