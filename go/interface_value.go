package main

import "fmt"

type A interface {
	X()
}

type B struct{}

func (b *B) X() {}

const (
	F, G int = iota, iota + 1
	_, _
	C, D
	E int = iota + 10
)

func main() {
	b := (A)((*B)(nil))
	fmt.Printf("%v\n", b == nil)
	fmt.Printf("%v\n", b == A(nil))
	fmt.Printf("%v\n", b == A((*B)(nil)))
	m := make(map[string]int)
	//  map的底层实现是hashtable 所以不能对key的value去地址
	//fmt.Println(&m["qcrao"]
	fmt.Println(m["qcrao"])
	fmt.Println(&m)
	fmt.Println(F, G, C, D, E)
}
