package main

import "fmt"

func main() {
	// GO : 모호하거나, 애매한 문법 금지
	// 후치연산자 허용(i++), 전치 연산자 비허용(++i) -> x
	// 증감연산 반환값 없음
	// 포인터(Pointer) 허용, 단 연산 비허용

	// example1
	sum, i := 0, 0
	for i <= 100 {
		sum += i
		i++
	}
	fmt.Println("ex1 : ", sum)

}
