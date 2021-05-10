package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func merge1(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for a != nil || b != nil {
			select {
			case v, ok := <-a:
				if !ok {
					a = nil
					fmt.Println("a is done")
					continue
				}
				c <- v
			case v, ok := <-b:
				if !ok {
					b = nil
					fmt.Println("b is done")
					continue
				}
				c <- v
			}
		}
	}()
	return c
}

func main() {
	fmt.Println("vim-go")
	a := asChan(1, 3, 5, 7, 9, 11, 13)
	b := asChan(2, 4, 6, 8, 10, 12, 14)
	c := merge1(a, b)
	for v := range c {
		fmt.Println(v)
	}
}
