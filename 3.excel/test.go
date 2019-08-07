package main

import "fmt"

func main() {
	//arr := [][]string
	arr := make([][]string, 10)

	fmt.Println(len(arr))

	var i int

	for i, _ := range arr{
		arr[i] = make([]string, 2)
		arr[i][0] ="123"
		arr[i][1] ="456"
		fmt.Println("i:%d", i)
	}

	//fmt.Println(arr)

	fmt.Println("==========")

	fmt.Println(i)

}