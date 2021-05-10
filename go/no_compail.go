package main

import "fmt"

type A struct {
}

func (a *A) X() {
	fmt.Println("test")
}

//func (a A) X() {
//	fmt.Println("test1")
//}

func main() {
	//	A{}.X()
	a := A{}
	a.X()
}
