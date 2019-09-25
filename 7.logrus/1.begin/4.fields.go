package main

import (
	"fmt"
	log2 "github.com/sirupsen/logrus"
	"log"
)

//原来的日志方式
func test1()  {
	log.Printf("Failed to send event %s to topic %s with key %", 11111, 222, 33333)
}

//logrus的方式
func test2()  {
	log2.WithFields(log2.Fields{
		"event": 111,
		"topic": 22,
		"key": 33,
	}).Fatal("Failed to send event")
}

/*
输出:
2019/09/24 11:58:32 Failed to send event %!s(int=11111) to topic %!s(int=222) with key %!(NOVERB)%!(EXTRA int=33333)
FATA[0000] Failed to send event                          event=111 key=33 topic=22
*/

func main() {
	test1()
	fmt.Println("=================================")
	test2()
}





