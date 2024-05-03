// 인터페이스 고급(3)
package main

import (
	"fmt"
)

func checkType(arg interface{}) {
	// arg.(Type)을 통해 현재 데이터형 반환
	switch arg.(type) {
	case bool:
		fmt.Println("This is a bool : ", arg)
	case int, int8, int16, int32, int64:
		fmt.Println("This is a int : ", arg)
	case float32, float64:
		fmt.Println("This is a float : ", arg)
	case string:
		fmt.Println("This is a string : ", arg)
	default:
		fmt.Println("What is this type? : ", arg)
	}
}

func main() {
	// 실제 타입 검사는 switch를 많이 쓴다.
	// 빈 인터페이스는 어떠한 자료형도 전달받을 수 있음으로, `타입체크`를 통해 형 변환 후 사용 가능

	// 예제1
	checkType(true)
	checkType(1)
	checkType(22.22)
	checkType("string")
	checkType([]int{})
}
