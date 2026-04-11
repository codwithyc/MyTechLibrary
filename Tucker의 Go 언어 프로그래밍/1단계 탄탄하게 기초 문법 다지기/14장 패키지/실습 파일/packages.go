package main

import (
	"fmt"
	"math"
)

const AppName = "Package Practice"

func init() {
	fmt.Println("init 함수가 main 함수보다 먼저 실행됩니다.")
}

func PrintMessage(message string) {
	fmt.Println(message)
}

func main() {
	PrintMessage(AppName)

	radius := 10.0
	area := math.Pi * radius * radius
	fmt.Printf("원의 넓이: %.2f\n", area)
}
