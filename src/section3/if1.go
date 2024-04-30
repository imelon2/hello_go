package main

import "fmt"

func main() {
	// 제어문(조건문)
	// IF문 : 반드시 Bool 검사 -> 자동 형 변환없어서 1,0 안됨

	var a int = 20
	b := 20
	if a >= 15 {
		fmt.Println("15 이상이닷!")
	}

	if b >= 25 {
		fmt.Println("25 이상이닷!")
	}

	if c := 40; c >= 40 {
		fmt.Println("35 이상")
	}

	// 짧은 변수는 정해진 함수 내에서만 실행 가능
	// fmt.Println(c)
}
