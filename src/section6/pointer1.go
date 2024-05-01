package main

import "fmt"

func main() {
	// Pointer (1)
	// Go : 포인터 지원, but 복잡하고 어려움 -> 위험성 존재
	// why : 변수의 지역성, 연솓된 메모미 참조 ..., 힙, 스택
	// Go는 주소의 값은 직접 변경 불가능(휴먼애러 방지)
	// * : 애스티리스크 사용
	// nil로 초기화 (nil == 0)

	// 예제1 - 포인터 주소값 할당 방법
	var a *int            // 방법 1
	var b *int = new(int) // 방법 2

	fmt.Println(a) //&
	fmt.Println(b)

	i := 7

	a = &i // & : 주소값 전달
	b = &i

	//   : 현재 값
	// & : 주소값 조회
	// * : 역 참조 조회 -> 현재 값이 가르키는 주소 값의 원본 값 조회
	fmt.Println("ex1 : ", a, &i)
	fmt.Println("ex1 : ", &a)
	fmt.Println("ex1 : ", *a)
	fmt.Println()
	fmt.Println("ex1 : ", b, &i)
	fmt.Println("ex1 : ", &b)
	fmt.Println("ex1 : ", *b)
	fmt.Println()
	var c = &i
	d := &i

	// 현재 a,b,c,d는 i를 가르키고 있으니, 참조(얕은복사)된 상태이다
	fmt.Println("ex1 : ", c, &i)
	fmt.Println("ex1 : ", &c)
	fmt.Println("ex1 : ", *c)
	fmt.Println()
	fmt.Println("ex1 : ", d, &i)
	fmt.Println("ex1 : ", &d)
	fmt.Println("ex1 : ", *d)
	fmt.Println()

	i = 99
	// 역참조의 값, 즉 포인터 i의 값을 변경
	// *d = 10

	fmt.Println("ex1 : ", a, &i)
	fmt.Println("ex1 : ", &a)
	fmt.Println("ex1 : ", *a)
	fmt.Println()
	fmt.Println("ex1 : ", b, &i)
	fmt.Println("ex1 : ", &b)
	fmt.Println("ex1 : ", *b)
	fmt.Println()
	fmt.Println("ex1 : ", c, &i)
	fmt.Println("ex1 : ", &c)
	fmt.Println("ex1 : ", *c)
	fmt.Println()
	fmt.Println("ex1 : ", d, &i)
	fmt.Println("ex1 : ", &d)
	fmt.Println("ex1 : ", *d)
	fmt.Println()
}
