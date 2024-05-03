package main

import "fmt"

func multiply4(x, y int) (r int) {
	r = x * y
	return r
}

func sum4(x, y int) (r int) {
	r = x + y
	return r
}

func main() {
	// 함수를 변수에 할당 - 재귀 가능

	// 예제1
	f := []func(int, int) int{multiply4, sum4}

	a := f[0](10, 10)
	b := f[1](10, 10)

	fmt.Println("Ex1 : ", a, b)

	// 예제2 - 변수에 할당
	var f1 func(int, int) int = multiply4 // 알아보기 쉬움
	f2 := sum4                            // 간단하지만 찾아가야함

	fmt.Println("Ex2 : ", f1(20, 20), f2(20, 20))

	// 예제3 - Map에 할당
	m := map[string]func(int, int) int{
		"mul_func": multiply4,
		"sum_func": sum4,
	}

	fmt.Println("Ex3 : ", m["mul_func"](30, 30))
	fmt.Println("Ex3 : ", m["sum_func"](30, 30))
}
