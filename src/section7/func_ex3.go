package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func prtHello(n int) {
	if n == 0 {
		return
	}
	fmt.Println("Ex2 : (", n, ")", "hi!!")
	prtHello(n - 1)
}

func main() {
	// 재귀 함수(Recursion)
	// 프로그램이 보기 쉽고, 코드 간결, 오류 수정이 용이 : 장점
	// 코드 이해하기 어렵고, 기억 공간을 많이 사용, 무한 루프 가능성 주의 : 단점

	// 예제1
	x := fact(7)
	fmt.Println("Ex1 : ", x)

	// 예제2
	prtHello(5)
}
