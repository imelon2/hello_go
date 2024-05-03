// 사용자 정의 타입(2)
package main

import "fmt"

type cnt int

func testConverT(i int) {
	fmt.Println("Ex3 : (Default Type) : ", i)
}

func testConverD(i cnt) {
	fmt.Println("Ex4 : (Default Type) : ", i)
}

func main() {
	// 기본 자료형 사용자 정의 타입

	// 예제1
	a := cnt(5) // 다른 함수가 있는거처럼 보여서
	fmt.Println("Ex1 : ", a)
	// 예제2
	var b cnt = 15 //  보통 이렇게 사용
	fmt.Println("Ex2 : ", b)

	// 예제3
	// testConverT(b) // type ERROR
	testConverT(int(b))
	testConverT(10)

	// 예제4
	testConverD(10)
	testConverD(cnt(10))
	testConverD(b)
}
