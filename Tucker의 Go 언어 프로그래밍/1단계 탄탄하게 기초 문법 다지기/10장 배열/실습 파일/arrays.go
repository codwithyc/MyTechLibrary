package main

import "fmt"

func main() {
	scores := [5]int{90, 85, 72, 100, 64}

	fmt.Println("=== 배열 출력 ===")
	fmt.Println(scores)
	fmt.Println("길이:", len(scores))

	fmt.Println("\n=== range 순회 ===")
	for index, score := range scores {
		fmt.Println(index, score)
	}

	fmt.Println("\n=== 배열 복사 ===")
	copied := scores
	copied[0] = 0
	fmt.Println("원본:", scores)
	fmt.Println("복사본:", copied)

	fmt.Println("\n=== 다중 배열 ===")
	board := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			fmt.Print(board[row][col], " ")
		}
		fmt.Println()
	}
}
