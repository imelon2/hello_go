package main

import (
	"fmt"
)

func main() {
	// 문자열 표현
	// 문자열 : UTF-8 인코딩
	// 실질적인 문자열은 Printf( %c)로 출력 할 수 있다.

	// 예제1
	var str1 string = "Golang" // = array
	var str2 string = "World"
	var str3 string = "고프로그래밍"

	fmt.Println("ex1 : ", str1[0], str1[1], str1[2], str1[3], str1[4], str1[5])
	fmt.Println("ex1 : ", str2[0], str1[1], str1[2], str1[3], str1[4])
	fmt.Println("ex1 : ", str3[0], str1[1], str1[2], str1[3], str1[4], str1[5])

	// 예제2
	fmt.Printf("ex2 : %c %c %c %c %c %c\n", str1[0], str1[1], str1[2], str1[3], str1[4], str1[5])
	fmt.Printf("ex2 : %c %c %c %c %c\n", str2[0], str2[1], str2[2], str2[3], str2[4])
	fmt.Printf("ex2 : %c %c %c %c %c %c\n", str3[0], str3[1], str3[2], str3[3], str3[4], str3[5])

	// 예제3
	conStr := []rune(str3)
	fmt.Printf("ex3 : %c %c %c %c %c %c\n", conStr[0], conStr[1], conStr[2], conStr[3], conStr[4], conStr[5])

	// 예제4
	for i, char := range str1 {
		fmt.Printf("ex3 : %c(%d)\n", char, i)
	}

	// 예제5
	for i, char := range str2 {
		fmt.Printf("ex4 : %c(%d)\n", char, i)
	}
}
