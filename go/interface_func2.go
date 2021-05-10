package main

import "fmt"

type Animaler interface {
	SetName(string)
	GetName() string
}

type Dog struct {
	Name string
}

func (d Dog) GetName() string {
	return d.Name
}

func (d *Dog) SetName(name string) {
	d.Name = name
}

type Cat struct {
	Name string
}

func (c Cat) SetName(name string) {
	c.Name = name
}

func (c Cat) GetName() string {
	return c.Name
}

func main() {
	fmt.Println("vim-go")

	// Dog实现Animal接口
	d := &Dog{Name: "二狗"}
	fmt.Println(d.GetName())
	d.SetName("阿龙")
	fmt.Println(d.GetName())
	var d1 Animaler = d
	fmt.Println(d1.GetName())
	d1.SetName("阿奇")
	fmt.Println(d1.GetName())

	// Cat实现Animal接口
	c := &Cat{Name: "花花"}
	var c1 Animaler = c
	fmt.Println(c1.GetName())

}
