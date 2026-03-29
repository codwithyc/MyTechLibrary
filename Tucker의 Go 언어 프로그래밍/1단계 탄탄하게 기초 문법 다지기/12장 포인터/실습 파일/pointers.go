package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func change(n *int) {
	*n = 100
}

func birthday(user *User) {
	user.Age++
}

func newUser(name string) *User {
	user := User{Name: name}
	return &user
}

func main() {
	age := 20
	p := &age

	fmt.Println("age:", age)
	fmt.Println("age 주소:", p)
	fmt.Println("포인터가 가리키는 값:", *p)

	*p = 30
	fmt.Println("변경된 age:", age)

	change(&age)
	fmt.Println("change 이후 age:", age)

	user := User{Name: "gopher", Age: 10}
	birthday(&user)
	fmt.Println("생일 이후:", user)

	created := newUser("new gopher")
	fmt.Println("생성된 사용자:", created)
}
