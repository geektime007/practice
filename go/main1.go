package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("vim-go")
	a := time.Time{}
	fmt.Printf("##:|%v|, %v", a, a.IsZero())
	fmt.Println()
	fmt.Printf("@@:%v\n", a.IsZero())

	fmt.Printf("##:%v", time.Now())
}
