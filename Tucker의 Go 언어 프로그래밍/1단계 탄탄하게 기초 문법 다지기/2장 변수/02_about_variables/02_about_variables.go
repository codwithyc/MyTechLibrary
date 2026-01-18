package main

import "fmt"

func main(){
	var age int = 20

	fmt.Println("이름 : ", "age")
	fmt.Println("값 : ",age)
	fmt.Println("타입 예시 : ", "int")
	fmt.Println("주소 : ", &age)

	var score int8 = 100
	var price float64 = 19.99

	fmt.Println("점수 : ", score)
	fmt.Println("가격 : ", price)
}