package main

import "fmt"

func main() {

	// switch 뒤 표현식(expression) 생략 가능
	// case 뒤 표현식(expression) 생략 가능
	// 자동 break 때문에 fallthrouth 존재
	// Type 분기 -> 값이 아닌 변수 Type으로 분기 가능

	// example 1
	a := -7
	switch {
	case a < 0:
		fmt.Println(a, "는 음수")
	case a == 0:
		fmt.Println(a, "는 0")
	case a > 0:
		fmt.Println(a, "는 양수")
	}

	// example 2
	switch b := 27; {
	case b < 0:
		fmt.Println(b, "는 음수")
	case b == 0:
		fmt.Println(b, "는 0")
	case b > 0:
		fmt.Println(b, "는 양수")
	}

	// example 3
	switch c := "go"; c {
	case "go":
		fmt.Println(c, "는 GO")
	case "java":
		fmt.Println(c, "는 JAVA")
	default:
		fmt.Println("일치하는 값 없음")
	}

	// example 4
	switch c := "go"; c + "lang" {
	case "golang":
		fmt.Println(c, "는 GO LANG")
	case "java":
		fmt.Println(c, "는 JAVA")
	default:
		fmt.Println("일치하는 값 없음")
	}

	// example 4
	switch i, j := 20, 30; {
	case i < j:
		fmt.Println("i는 j보다 작다")
	case i == j:
		fmt.Println("i는 j와 같다")
	case i > j:
		fmt.Println("i는 j보다 다")

	}
}
