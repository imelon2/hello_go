package main

import "fmt"

func main() {
	// 예제1
	cnt := increaseCnt()

	fmt.Println("Ex1 : ", cnt())
	fmt.Println("Ex1 : ", cnt())
	fmt.Println("Ex1 : ", cnt())
	fmt.Println("Ex1 : ", cnt())
	fmt.Println("Ex1 : ", cnt())
	fmt.Println("Ex1 : ", cnt())
}

func increaseCnt() func() int {
	n := 0 // 지역변수 스냅샷 - 함수 종료되도 변수(선언된 상태로) 유지
	return func() int {
		n += 1
		return n
	}
}
