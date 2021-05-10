package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	err "github.com/geektime007/errors"
	"github.com/geektime007/practice/go/error/errors"
)

func PrintMsg() *err.Error {
	fmt.Println("helloworld")
	return errors.Unknown
}

func main() {
	fmt.Println("vim-go")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.GET("/helloworld", func(c *gin.Context) {
		ret := PrintMsg()
		fmt.Println(ret)
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.Run()
}
