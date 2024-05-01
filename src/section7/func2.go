package main

import (
	"fmt"
)

func sum1(i int, f func(int, int)) {
	f(i, 10)
}

// func add(a int, b int) {
func add(a, b int) {
	fmt.Println("ex1 : ", a+b)
}

func multi_value(i int) {
	i = i * 10
}

func multi_reference(i *int) {
	*i = *i * 10
}
func main() {

	// 예제1 - 함수(콜백) 전달,
	sum1(100, add)

	// 예제2 - 값 전달(call by value)
	a := 100
	multi_value(a)
	fmt.Println("ex2 : ", a)

	// 예제3 - 참조 전달(call by reference)
	multi_reference(&a)
	fmt.Println("ex3 : ", a)
}
