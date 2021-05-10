package main

import "fmt"

type People struct {
	Name string
}

func (p *People) String() string {
	fmt.Sprintf("print: 1")
	return fmt.Sprintf("print: %v", p.Name)
}

func main() {
	fmt.Println("vim-go")
	p := &People{Name: "nihao"}
	fmt.Println(p)
}
