package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type A struct {
	s string
}

func foo(s string) *A {
	//a := new(A)
	a := &A{}
	a.s = s
	return a
}

func main() {
	fmt.Println("vim-go")
	a := foo("hello")
	fmt.Println(unsafe.Pointer(a))
	fmt.Println(unsafe.Pointer(&a))
	fmt.Println((*reflect.StringHeader)(unsafe.Pointer(a)))
	fmt.Println((*reflect.StringHeader)(unsafe.Pointer(&a)))

	fmt.Println(*(*reflect.StringHeader)(unsafe.Pointer(a)))
	fmt.Println(*(*reflect.StringHeader)(unsafe.Pointer(&a)))

	b := a.s + "world"
	c := b + "!"
	fmt.Println(c)
}
