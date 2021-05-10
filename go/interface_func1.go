package main

import "fmt"

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {}

func live() People {
	var stu *Student
	fmt.Println("stu = ", stu)
	return stu
}

func main() {
	fmt.Println("vim-go")

	t := live()
	if t == nil {
		fmt.Println("AAAAAAAAA")
	} else {
		fmt.Printf("BBBBBBBBBB, t= %v , live= %v\n", t, live())
	}
}
