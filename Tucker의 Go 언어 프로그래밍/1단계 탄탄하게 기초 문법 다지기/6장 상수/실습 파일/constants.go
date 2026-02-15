package main

import "fmt"

const pi = 3.141592
const maxRetry = 3

const (
	statusReady = iota
	statusRunning
	statusDone
	statusFailed
)

func statusName(status int) string {
	switch status {
	case statusReady:
		return "ready"
	case statusRunning:
		return "running"
	case statusDone:
		return "done"
	case statusFailed:
		return "failed"
	default:
		return "unknown"
	}
}

func main() {
	radius := 10.0
	area := radius * radius * pi

	fmt.Printf("원의 넓이: %.2f\n", area)
	fmt.Println("최대 재시도 횟수:", maxRetry)
	fmt.Println("현재 상태:", statusName(statusRunning))

	const number = 10
	var intValue int = number
	var floatValue float64 = number

	fmt.Println("타입 없는 상수 사용:", intValue, floatValue)
}
