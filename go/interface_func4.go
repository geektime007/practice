package main

import "fmt"

type Address struct {
	Name  string
	Phone int
}

func main() {
	fmt.Println("vim-go")
	var userInfo = make(map[string]interface{})
	userInfo["userName"] = "张三"
	userInfo["age"] = 20
	userInfo["hobby"] = []string{"吃饭", "睡觉"}
	fmt.Println(userInfo["age"])
	fmt.Println(userInfo["hobby"])
	//fmt.Println(userInfo["hobby"][1])

	hobby2, _ := userInfo["hobby"].([]string)
	fmt.Println(hobby2[1]) // 睡觉

	var address = Address{
		Name:  "李四",
		Phone: 13912345678,
	}

	fmt.Println(address.Name)

	userInfo["address"] = address
	fmt.Println(userInfo["address"])

	//var name = userInfo["address"].Name
	// type interface {} is interface with no methods空接口中没有name方法

	address2, _ := userInfo["address"].(Address)
	fmt.Println(address2.Name)

}
