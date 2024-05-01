package main

import (
	"fmt"
)

func multiply(x, y int) (int, int) {
	return x * 10, y * 10
}

func arrayMultiply(a, b, c, d, e int) (int, int, int, int, int) {
	return a * 1, b * 2, c * 3, b * 4, e * 5
}

func main() {
	// 다중 값 변환
	// 예제1
	a, b := multiply(10, 5)
	_, c := multiply(10, 5)
	d, _ := multiply(10, 5)

	fmt.Println("ex1 : ", a, b, c, d)

	// 예제2
	x1, x2, x3, x4, x5 := arrayMultiply(1, 2, 3, 4, 5)
	y1, _, _, y4, _ := arrayMultiply(1, 2, 3, 4, 5)

	fmt.Println("ex2 : ", x1, x2, x3, x4, x5)
	fmt.Println("ex2 : ", y1, y4)
}
