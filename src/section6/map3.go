package main

import "fmt"

func main() {
	// Map 값 변경 및 삭제

	// 예제1
	map1 := map[string]string{
		"daum":   "daum.net",
		"naver":  "naver.com",
		"google": "google.com",
		"home1":  "test1.com",
	}
	fmt.Println("ex1 : ", map1)
	fmt.Println()

	map1["home2"] = "test2.com" // insert
	fmt.Println("ex1 : ", map1)
	fmt.Println()

	map1["home2"] = "test2-2.com" //  update
	fmt.Println("ex1 : ", map1)
	fmt.Println()

	// 예제2 - 삭제
	delete(map1, "home2")
	fmt.Println("ex2 : ", map1) // delete
	fmt.Println()
}
