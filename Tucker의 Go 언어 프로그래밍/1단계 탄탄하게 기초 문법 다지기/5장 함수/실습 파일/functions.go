package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}

	return a / b, true
}

func isEven(number int) bool {
	return number%2 == 0
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}

	return n * factorial(n-1)
}

func change(n int) {
	n = 100
}

func main() {
	fmt.Println("add:", add(10, 20))

	result, ok := divide(10, 2)
	if ok {
		fmt.Println("divide:", result)
	}

	number := 8
	fmt.Println("isEven:", isEven(number))
	fmt.Println("factorial:", factorial(5))

	age := 20
	change(age)
	fmt.Println("change 이후 age:", age)
}
