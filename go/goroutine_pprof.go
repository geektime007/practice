package main

import (
	"fmt"
	"net/http"

	_ "net/http/pprof"
)

func main() {
	fmt.Println("vim-go")
	http.ListenAndServe("localhost:6060", nil)
}
