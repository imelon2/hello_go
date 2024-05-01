package main

import "fmt"

func main() {
	// 예제1
	arr1 := [5]int{1, 10, 100, 1000, 10000}

	for i := 0; i < len(arr1); i++ {
		fmt.Println("ex1 : ", arr1[i])
	}

	// 예제2 - 가장 많이 사용
	arr2 := [5]int{7, 77, 777, 7777, 77777}
	for i, v := range arr2 {
		fmt.Println("ex2 : ", i, v)
	}

	// 예제3 - 배열의 값을 순서대로 사용
	for _, v := range arr2 {
		fmt.Println("ex3 : ", v)
	}

	// 예제4 - for 첫번째 변수는 index 고정!!
	for v := range arr2 {
		fmt.Println("ex4 : ", v)
	}
}
