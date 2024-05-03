// 인터페이스 고급(1)
package main

import "fmt"

func main() {
	// 인터페이스 활용 - 빈 인터페이스
	// 빈 인터페이스 - 함수 매개변수, 리턴 값, 구조체 필드 등으로 사용 가능 -> 어떤 타입으로도 변환 가능
	// Object type(동적)과 같음

	var a interface{}
	printContents(a)

	a = 4.5
	printContents(a)
	a = true
	printContents(a)
	a = "만능 인터페이스"
	printContents(a)
}

func printContents(v interface{}) {
	fmt.Printf("Type : (%T)", v)
	fmt.Println("Ex1 : ", v) // 원본 타입
}
