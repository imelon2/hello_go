package main

import (
	"fmt"
)

// 데이터 타입 : 숫자형
// 정수, 실수, 복소수
// 32bit, 64bit, unsigned(양수)
// 8진수, 16진수, 10진수
func main() {
	var num1 int = 17
	var num2 int = -17
	var num3 int = 0631 // 8진수
	var num4 int = 0xff // 16진수

	fmt.Println("num1 : ", num1)
	fmt.Println("num2 : ", num2)
	fmt.Println("num3 : ", num3)
	fmt.Println("num4 : ", num4)
}
