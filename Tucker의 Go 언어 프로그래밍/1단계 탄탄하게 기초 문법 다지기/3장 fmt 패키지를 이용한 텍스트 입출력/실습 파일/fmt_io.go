package main

import "fmt"

func main() {
	name := "Go"
	version := 1.22

	fmt.Println("=== 출력 예제 ===")
	fmt.Print("Hello ")
	fmt.Println(name)
	fmt.Printf("%s version %.2f\n", name, version)

	var userName string
	var age int

	fmt.Println("\n=== 입력 예제 ===")
	fmt.Print("이름과 나이를 입력하세요: ")

	count, err := fmt.Scan(&userName, &age)
	if err != nil {
		fmt.Println("입력 오류:", err)
		return
	}

	fmt.Println("읽은 값 개수:", count)
	fmt.Printf("이름: %s, 나이: %d\n", userName, age)
}
