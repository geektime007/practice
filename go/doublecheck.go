package main

import (
	"fmt"
	"sync"
)

type Once struct {
	m    sync.Mutex
	done uint32
}

func (o *Once) Do(f func()) {
	if o.done == 1 {
		return
	}

	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		o.done = 1
		f()
	}
}



func main() {
	fmt.Println("vim-go")
	one := &Once{
		m: sync.Mutex{}
	}

	for i:=0; i< 100; i++ {
		go
	}
}
