package main

import "fmt"

func divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}

	return a / b, true
}

func main() {
	score := 85

	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}

	age := 22
	hasTicket := true

	if age >= 20 && hasTicket {
		fmt.Println("입장할 수 있습니다.")
	}

	if result, ok := divide(10, 2); ok {
		fmt.Println("나눗셈 결과:", result)
	} else {
		fmt.Println("나눗셈에 실패했습니다.")
	}
}
