package main

import "fmt"

func main() {
	// 익명함수 사용할 경우, 함수를 변수에 할당해서 사용 가능
	// 함수 안에서 함수를 선언 및 정의 가능 -> 이때 외부 함수에 선언된 변수에 접근 가능한 함수
	// 함수가 선언되는 순간에, 함수가 실행될 때, 실체의 외부 변수에 접근하기 위한 스냅샷(객체)
	// 이전의 값을 유지하기 위해 사용

	// 예제1 - Closure 사용안하는 경우
	multiply := func(x, y int) int {
		return x * y
	}

	r1 := multiply(5, 10)
	fmt.Println("Ex1 : ", r1)

	// 예제2
	m, n := 5, 10 // 스냅샷
	sum := func(c int) int {
		// m, n은 내부가 아닌 외부 데이터를 갖고옴
		return m + n + c
	}

	r2 := sum(10)
	fmt.Println("Ex2 : ", r2)
}
