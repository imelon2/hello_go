package main

import (
	"fmt"
)

func main() {
	str1 := "Golang"
	str2 := "World"

	fmt.Println("ex1 : ", str1 == str2)
	fmt.Println("ex2 : ", str1 != str2)

	// 문자열 -> 아스키 코드에 대한 사전식 비교
	fmt.Println("ex3 : ", str1 > str2) // false
	fmt.Println("ex4 : ", str1 < str2) // true
}
