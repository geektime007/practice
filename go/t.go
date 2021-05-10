package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type People struct {
	Name string `json:"name"`
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(time.Now().Unix())
	//fmt.Printf("## 5*time.Second=%v", 5*time.Second)

	str := "Fan jiao long"
	count := strings.Count(str, " ")
	fmt.Printf("count=%v\n", count)

	js := `{"name":"11"}`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("people: ", p)
	fmt.Printf("people: %v", p)
}
