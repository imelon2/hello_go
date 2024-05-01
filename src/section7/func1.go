package main

import (
	"fmt"
	"strconv"
)

func helloGolang() {
	fmt.Println("ex1 : Hello Golang!!")
}

func say_one(m string) {
	fmt.Println("ex2 : ", m)
}

func sum(x int, y int) int {
	return x + y
}

func main() {
	// 함수
	// 함수 선언 : func

	// 예제1
	helloGolang()

	// 예제2
	say_one("Hello Say One !!")

	// 예제3
	fmt.Println("ex3 : ", sum(5, 2))
	// strconv.Itoa() -> string으로 type 변환
	fmt.Println("ex3 : ", strconv.Itoa(sum(10, 5)))
}
