package main

import (
	"fmt"
	"time"
)

var generateTaskIdSince int64 = 1546300800

func main() {
	fmt.Println("vim-go")
	now := time.Now().Unix()
	fmt.Println(now, generateTaskIdSince)
	fmt.Printf("%b,\n%b\n", now, generateTaskIdSince)

	fmt.Printf("%b\n", now-generateTaskIdSince)
	timePart := (now - generateTaskIdSince) & 0b111111111111111111111111111111
	fmt.Printf("%b\n", timePart)
	fmt.Println(timePart)
	a := 1
	fmt.Printf("%b\n,%b\n,%v\n", a, a^0b1, ^a)
	fmt.Printf("%b\n", a<<1)
	fmt.Println(1^1, ^1, ^0)
}
