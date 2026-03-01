package main

import "fmt"

const (
	statusReady = iota
	statusRunning
	statusDone
	statusFailed
)

func main() {
	month := 1

	switch month {
	case 12, 1, 2:
		fmt.Println("겨울")
	case 3, 4, 5:
		fmt.Println("봄")
	case 6, 7, 8:
		fmt.Println("여름")
	case 9, 10, 11:
		fmt.Println("가을")
	default:
		fmt.Println("잘못된 월입니다.")
	}

	score := 87

	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	default:
		fmt.Println("D")
	}

	status := statusRunning

	switch status {
	case statusReady:
		fmt.Println("준비 중입니다.")
	case statusRunning:
		fmt.Println("실행 중입니다.")
	case statusDone:
		fmt.Println("완료되었습니다.")
	case statusFailed:
		fmt.Println("실패했습니다.")
	}
}
