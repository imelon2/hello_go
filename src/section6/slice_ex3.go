package main

import "fmt"

func main() {
	// 슬라이스 복사
	// copy(복사대상, 원본)
	// make로 공간 할당 후 copy 사용 가능
	// 복사된 슬라이스 값 변경은 원본에 영향 없음 -> 깊은 복사

	// 예제1 (복사)
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := make([]int, 5)
	slice3 := []int{}

	fmt.Println("ex1 : ", slice1)
	fmt.Println("ex1 : ", slice2)
	fmt.Println("ex1 : ", slice3)
	fmt.Println()

	// `정해진 용량 만큼만 복사`되어 = [1 2 3 4 5]
	copy(slice2, slice1)
	// `정해진 용량 만큼만 복사`되어 = []
	copy(slice3, slice1)

	fmt.Println("ex1 : ", slice1) // 원본은 변화 없음
	fmt.Println("ex1 : ", slice2)
	fmt.Println("ex1 : ", slice3)
	fmt.Println()

	// 예제2
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5)

	copy(b, a)

	b[0] = 7
	b[4] = 10

	fmt.Println("ex2 : ", a) // 원본은 변화 없음
	fmt.Println("ex2 : ", b)
	fmt.Println()

	// 예제3 - 슬라이스 추출은 참조복사(얕은복사)!!!!
	c := [5]int{1, 2, 3, 4, 5}
	d := c[0:3]

	d[1] = 7

	// 원본도 바뀜 -> 슬라이스 추출은 참조 복사(얕은복사)!!
	fmt.Println("ex3 : ", c)
	fmt.Println("ex3 : ", d)
	fmt.Println()

	// 예제4 - 용량 지정 슬라이스 추출
	e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f := e[0:5:7] // [i:j:용량]

	fmt.Println("ex4 : ", len(f), cap(f), f)
}
