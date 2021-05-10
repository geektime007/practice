package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	// 数组只能与相同纬度长度以及类型的其他数组比较
	fmt.Println([...]string{"1"} == [...]string{"1"})

	// 切片之间不能直接比较
	//fmt.Println([]string{"1"} == []string{"1"})
}
