package main

import (
	"fmt"
)

func main() {
	// 길이 고정x -> 동적으로 크기가 늘어난다(가변적)
	// 레퍼런스(참조 값) 타입

	// 슬라이스 (길이 & 용량) 크기가 동적으로 할당

	// 2가지 선언 방법
	// 1. 배열처럼 선언

	// 예제1
	// 2. make 함수 : make(자료형, 길이, 용량(생략시 길이))
	var slice1 []int
	slice2 := []int{}
	slice3 := []int{1, 2, 3, 4, 5}
	slice4 := [][]int{
		{1, 1, 1, 1, 1},
		{2, 2, 2, 2, 2},
		{3, 3, 3, 3, 3},
		{4, 4, 4, 4, 4},
		{5, 5, 5, 5, 5},
	}

	// slice는 가변형으로, 선언시 0으로 초기화되지 않음. 때문에 초기 값이 없으면 `slice2[0] = 1`같이 값 변경은 ERROR가됨
	// slice2[0] = 1 // ERROR
	slice3[4] = 10

	fmt.Printf("%-5T %d %v\n", slice1, len(slice1), slice1)
	fmt.Printf("%-5T,%d %v\n", slice2, len(slice2), slice2)
	fmt.Printf("%-5T,%d %v\n", slice3, len(slice3), slice3)
	fmt.Printf("%-5T,%d %v\n", slice4, len(slice4), slice4)

	// 예제2
	// ake 함수 : make(자료형, 길이, 용량(생략시 =길이))
	// Go는 입력한 용량만큼 미리 메모리를 할당시킨다.(용량과길이.png 참고)
	// 하지만 슬라이스는 가변적이기에, 입력한 용량을 초과하면 메모리를 `재할당`하며,이때 성능 저하가 발생한다.
	var slice5 []int = make([]int, 5, 10)
	var slice6 = make([]int, 5)
	slice7 := make([]int, 5, 100)
	slice8 := make([]int, 5)
	fmt.Println(slice5)

	slice6[2] = 7

	// *예제1과 다르게 예제2는 0으로 초기화된 모습
	fmt.Printf("%-5T %d %v\n", slice5, len(slice5), slice5)
	fmt.Printf("%-5T,%d %v\n", slice6, len(slice6), slice6)
	fmt.Printf("%-5T,%d %v\n", slice7, len(slice7), slice7)
	fmt.Printf("%-5T,%d %v\n", slice8, len(slice8), slice8)

	// 예제3
	var slice9 []int // nil 슬라이스(길이와 용량 모두 0)
	if slice9 == nil {
		fmt.Println("This is nil")
	}
}
