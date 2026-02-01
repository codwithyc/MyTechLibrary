package main

import "fmt"

func main() {
	a := 10
	b := 3

	fmt.Println("=== 산술 연산자 ===")
	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a * b =", a*b)
	fmt.Println("a / b =", a/b)
	fmt.Println("a % b =", a%b)

	fmt.Println("\n=== 비교 연산자 ===")
	fmt.Println("a == b:", a == b)
	fmt.Println("a != b:", a != b)
	fmt.Println("a > b:", a > b)

	fmt.Println("\n=== 실수 오차 ===")
	x := 0.1
	y := 0.2
	sum := x + y
	fmt.Printf("0.1 + 0.2 = %.17f\n", sum)
	fmt.Println("sum == 0.3:", sum == 0.3)

	fmt.Println("\n=== 논리 연산자 ===")
	age := 22
	hasTicket := true
	fmt.Println("입장 가능:", age >= 20 && hasTicket)

	fmt.Println("\n=== 대입 연산자 ===")
	count := 10
	count += 5
	count++
	fmt.Println("count:", count)
}
