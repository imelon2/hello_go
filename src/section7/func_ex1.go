package main

import (
	"fmt"
)

func multiply3(n ...int) int {
	tot := 1
	for _, value := range n {
		tot *= value
	}

	return tot
}

func sum3(n ...int) int {
	tot := 1
	for _, value := range n {
		tot += value
	}

	return tot
}

func prtWord(msg ...string) {
	for _, value := range msg {
		fmt.Println("By prtWord : ", value)
	}
}

func main() {
	// 함수 고급
	// 가변 인자(매개 변수 개수가 동적으로 변할 때)

	// 예제1
	x := multiply3(5, 6, 7, 8, 9, 10)
	y := sum3(5, 6, 7, 8)

	fmt.Println("Ex1 : ", x)
	fmt.Println("Ex1 : ", y)

	// 예제2
	prtWord("a", "b", "c", "d", "e")

	// 예제3
	a := []int{5, 6, 7, 8, 9, 10}

	// m:=multiply3(a) // ERROR
	m := multiply3(a...)
	n := sum3(a...)

	fmt.Println("Ex3 : ", m)
	fmt.Println("Ex3 : ", n)
}
