package main

import (
	"fmt"
)

func multiply1(x, y int) (r1 int, r2 int) {
	r1 = x * 10
	r2 = y * 20
	return r1, r2
}

func main() {
	// 예제1
	// return 값 변수 사용
	a, b := multiply1(10, 5)
	fmt.Println("Ex1 : ", a, b)

}
