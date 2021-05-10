package main

import "fmt"

type Student struct {
	Age int
}

func main() {
	fmt.Println("vim-go")

	kv := map[string]Student{"menglu": {Age: 21}}
	//kv["menglu"].Age = 22
	fmt.Printf("%T\n", kv["menglu"])
	s := []Student{{Age: 21}}
	s[0].Age = 22
	fmt.Println(kv, s)
	fmt.Println(1 << 5)
}
