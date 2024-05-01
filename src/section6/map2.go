package main

import "fmt"

func main() {
	// Map 조회 및 순회(Iterator)

	// 예제1
	map1 := map[string]string{
		"daum":   "daum.net",
		"naver":  "naver.com",
		"google": "google.com",
	}

	fmt.Println("Ex1 : ", map1["google"])
	fmt.Println("Ex1 : ", map1["daum"])

	// 예제2 순회(Iterator) - 순서없음으로 랜덤
	for _, v := range map1 {
		fmt.Println("ex2 : ", v)
	}

	fmt.Println()

}
