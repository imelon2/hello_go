package main

import "fmt"

func main() {
	// 익명함수
	// 즉시 실행(선언과 동시에)

	// 예제1
	func() {
		fmt.Println("Ex1 : 어나니머스 함수")
	}()

	// 예제2
	msg := "Hello Golang!"

	func(m string) {
		fmt.Println("Ex2 : ", m)
	}(msg)

	// 예제3
	func(x, y int) {
		fmt.Println("Ex3 : ", x+y)
	}(10, 20)

	// 예제4
	r := func(x, y int) int {
		return x * y
	}(10, 20)

	fmt.Println("Ex4 : ", r)
}
