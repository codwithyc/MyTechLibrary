package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 10
	var result float64 = float64(num)

	fmt.Println(num)
	fmt.Println(result)

	var pi float64 = 3.9
	var intPi int = int(pi)

	fmt.Println(pi)
	fmt.Println(intPi)

	text := strconv.Itoa(num)
	fmt.Println(text)

	value, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("변환 실패:", err)
		return
	}

	fmt.Println(value)
}