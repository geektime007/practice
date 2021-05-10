package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("vim-go")
	sigs := []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT}
	c := make(chan os.Signal, 1)
	d := make(chan bool)
	signal.Notify(c, sigs...)
	go func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Printf("sleeping...\t")
			select {
			case <-d:
				return
			default:
				fmt.Printf("default next...\n")
			}
		}
	}()
	//go func() {
	for {
		select {
		case <-c:
			fmt.Println("Receive single exiting.")
			d <- true
			return
		}
	}
	//}()

	//select {}
}
