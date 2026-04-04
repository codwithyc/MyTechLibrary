package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	text := "고 Go"

	fmt.Println("문자열:", text)
	fmt.Println("바이트 길이:", len(text))
	fmt.Println("문자 개수:", utf8.RuneCountInString(text))

	fmt.Println("\n=== 바이트 순회 ===")
	for i := 0; i < len(text); i++ {
		fmt.Printf("index: %d, byte: %d\n", i, text[i])
	}

	fmt.Println("\n=== rune 순회 ===")
	for index, r := range text {
		fmt.Printf("index: %d, rune: %c\n", index, r)
	}

	message := "Hello" + " " + "Go"
	fmt.Println("\n문자열 합치기:", message)

	var builder strings.Builder
	for i := 0; i < 3; i++ {
		builder.WriteString("Go ")
	}
	fmt.Println("Builder 결과:", builder.String())

	runes := []rune("hello")
	runes[0] = 'H'
	fmt.Println("수정한 문자열:", string(runes))
}
