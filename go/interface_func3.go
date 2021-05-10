package main

import "fmt"

type Ainterface interface {
	SetName(string)
}

type Binterface interface {
	GetName() string
}

// 接口嵌套
type Animaler interface {
	Ainterface
	Binterface
}

type Dog struct {
	Name string
}

func (d *Dog) SetName(name string) {
	d.Name = name
}

func (d *Dog) GetName() string {
	return d.Name
}

func main() {
	fmt.Println("vim-go")
	// 让dog实现Animaler接口
	d := &Dog{Name: "阿龙"}
	var d1 Animaler = d
	d1.SetName("小阿龙")
	fmt.Println(d1.GetName())
}
