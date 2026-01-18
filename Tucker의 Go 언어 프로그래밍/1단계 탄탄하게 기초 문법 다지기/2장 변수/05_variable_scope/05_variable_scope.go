package main

import "fmt"

var globalMessage string = "전역 변수"

func main() {
	var localMessage string = "지역 변수"

	fmt.Println(globalMessage)
	fmt.Println(localMessage)

	if true {
		blockMessage := "블록 변수"
		fmt.Println(blockMessage)
	}

	printGlobal()
}

func printGlobal() {
	fmt.Println(globalMessage)
}