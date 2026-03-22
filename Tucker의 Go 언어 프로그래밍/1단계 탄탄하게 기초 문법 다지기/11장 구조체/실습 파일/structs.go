package main

import "fmt"

type Address struct {
	City   string
	Street string
}

type User struct {
	Name    string
	Age     int
	Address Address
}

type Character struct {
	Name  string
	Level int
	HP    int
}

func main() {
	user := User{
		Name: "gopher",
		Age:  10,
		Address: Address{
			City:   "Seoul",
			Street: "Go-ro",
		},
	}

	fmt.Println("사용자:", user)
	fmt.Println("도시:", user.Address.City)

	user.Age = 11
	fmt.Println("변경된 나이:", user.Age)

	character := Character{
		Name:  "Gopher Knight",
		Level: 3,
		HP:    100,
	}

	fmt.Printf("캐릭터: %+v\n", character)
}
