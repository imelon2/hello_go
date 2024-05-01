package main

import (
	"fmt"
	"sort"
)

func main() {
	// 슬라이스 추출 및 정령
	// slice[i:j] -> i ~ j-1 까지 추출
	// slice[i:] -> i ~ 마지막 까지 추출
	// slice[:j] -> 처음 ~ j 까지 추출
	// slice[:] -> 처음 ~ 마지막 까지 추출

	// 예제1(추출)
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("ex1 : ", slice1[:])
	fmt.Println("ex1 : ", slice1[0:])
	fmt.Println("ex1 : ", slice1[:5])
	fmt.Println("ex1 : ", slice1[0:len(slice1)])
	fmt.Println("ex1 : ", slice1[1:3])
	fmt.Println()

	// 예제2(정렬)
	// sort 패키지
	slice2 := []int{1, 100, 6, 12, 59}
	slice3 := []string{"b", "c", "a", "e"}

	//  sort.IntsAreSorted() : Int 렬 확인 함수
	fmt.Println("ex2 : ", sort.IntsAreSorted(slice2))
	// sort.Ints() : Int 정렬 함수
	sort.Ints(slice2)
	fmt.Println("ex2 : ", slice2)

	//  sort.StringsAreSorted() : String 정렬 확인 함수
	fmt.Println("ex2 : ", sort.StringsAreSorted(slice3))

	// sort.Strings() : 문자열 정렬 함수
	sort.Strings(slice3)
	fmt.Println("ex2 : ", slice3)
	fmt.Println("ex2 : ", sort.StringsAreSorted(slice3))

}
