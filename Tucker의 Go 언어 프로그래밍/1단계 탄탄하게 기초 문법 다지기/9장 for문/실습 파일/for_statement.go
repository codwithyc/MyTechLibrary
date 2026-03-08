package main

import "fmt"

func main() {
	fmt.Println("=== 기본 for문 ===")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("\n=== continue ===")
	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("\n=== 조건만 있는 for문 ===")
	count := 3
	for count > 0 {
		fmt.Println(count)
		count--
	}

	fmt.Println("\n=== 중첩 for문 ===")
	for dan := 2; dan <= 3; dan++ {
		for num := 1; num <= 3; num++ {
			fmt.Printf("%d x %d = %d\n", dan, num, dan*num)
		}
	}

	fmt.Println("\n=== 레이블 break ===")
Outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break Outer
			}
			fmt.Println(i, j)
		}
	}
}
