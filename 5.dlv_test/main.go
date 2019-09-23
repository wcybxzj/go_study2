package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func test()  {
	for i:=0; i<1; i++ {
		//time.Sleep(time.Duration(2)*time.Second)
		fmt.Print("11111")
	}
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		test()
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}