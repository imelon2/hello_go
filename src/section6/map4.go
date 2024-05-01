package main

import "fmt"

func main() {
	// Map 조회할 경우 `주의` 할 점

	// 예제1
	map1 := map[string]int{
		"apple":  15,
		"banana": 40,
		"orange": 23,
	}

	value1, ok1 := map1["apple"] // 있는 값
	value2 := map1["kiwi"]       // 없는 값 -> 기본 초기화 값
	value3, ok := map1["kiwi"]   // 없는 값 -> 기본 초기화 값,존재 여부(bool)

	fmt.Println("ex1 : ", value1, ok1)
	fmt.Println("ex1 : ", value2)
	fmt.Println("ex1 : ", value3, ok)
	fmt.Println()

	// 예제2 - 존재 여부(bool) 활용방법
	if value, ok := map1["kiwi"]; ok {
		fmt.Println("ex2 : ", value)
	} else {
		fmt.Println("ex2 : ", "Kiwi is not exist!!")
	}

	if value, ok := map1["apple"]; ok {
		fmt.Println("ex2 : ", value)
	} else {
		fmt.Println("ex2 : ", "apple is not exist!!")
	}

	if _, ok := map1["apple"]; ok {
		fmt.Println("ex2 : ", "apple 있어요!!")
	} else {
		fmt.Println("ex2 : ", "apple is not exist!!")
	}
}
