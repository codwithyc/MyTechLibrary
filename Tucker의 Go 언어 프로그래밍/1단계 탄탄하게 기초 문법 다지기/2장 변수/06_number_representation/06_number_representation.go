package main

import "fmt"

func main() {
	var decimal int = 10
	var binary int = 0b1010
	var octal int = 0o12
	var hex int = 0xA

	var floatNum float64 = 3.14

	fmt.Println("10진수:", decimal)
	fmt.Println("2진수:", binary)
	fmt.Println("8진수:", octal)
	fmt.Println("16진수:", hex)
	fmt.Println("실수:", floatNum)

	fmt.Println("실수 정밀도 확인:", 0.1+0.2)
}