package main

import (
	"fmt"
	"sort"
)

func main() {
	// 기본 생성과 길이, 용량 확인
	numbers := make([]int, 3, 5)
	numbers[0], numbers[1], numbers[2] = 10, 20, 30

	fmt.Println("numbers:", numbers)
	fmt.Println("len(numbers):", len(numbers))
	fmt.Println("cap(numbers):", cap(numbers))

	// 같은 기반 배열을 공유하는 슬라이스
	part := numbers[:2]
	part[1] = 999

	fmt.Println("part:", part)
	fmt.Println("shared numbers:", numbers)

	// append와 용량 증가 확인
	numbers = append(numbers, 40, 50)
	fmt.Println("after append:", numbers)
	fmt.Println("len/cap:", len(numbers), cap(numbers))

	// copy로 독립 복사본 만들기
	cloned := make([]int, len(numbers))
	copy(cloned, numbers)
	cloned[0] = -1

	fmt.Println("original:", numbers)
	fmt.Println("cloned:", cloned)

	// 슬라이싱 활용
	queue := []int{1, 2, 3, 4}
	queue = queue[1:]
	fmt.Println("queue after pop front:", queue)

	stack := []int{1, 2, 3, 4}
	stack = stack[:len(stack)-1]
	fmt.Println("stack after pop back:", stack)

	// 정렬
	values := []int{30, 10, 50, 20, 40}
	sort.Ints(values)
	fmt.Println("ascending:", values)

	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	fmt.Println("descending:", values)
}
