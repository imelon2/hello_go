package main

import "fmt"

func rptc(n *int) {
	fmt.Println("By rptc() -> `n` pointer : ", n)
	*n = 77
}

func vptc(n int) {
	n = 77 // 깊은 복사된 `n`은 함수 종료 후 소멸
}

func main() {
	// 포인터 값 전달
	// 함수, Method 호출 시, 매개변수 값을 `깊은 복사 전달`
	// 함수 종료전 까지 소멸되지 않아 성능 구령
	// 특히 크기가 큰 배열인 경우, 매번 복사 시, 시스템 부담함수 종료 -> 포인터로 해결

	// 예제1
	var a int = 10
	var b int = 10

	fmt.Println("Ex1 Before : ", a)
	fmt.Println("Ex1 Before : ", b)
	fmt.Println()

	rptc(&a) // Pointer 전달 -> 얕은 복사
	vptc(b)  //  -> 깊은 복사

	fmt.Println("Ex1 After : ", a)
	fmt.Println("Ex1 After : ", b)
	fmt.Println()

}
