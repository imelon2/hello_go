// 인터페이스 고급(2)
package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 인터페이스 활용 - type assertion(타입변환)
	// 실행(런타임) 중에는 인터페이스에 할당한 변수는 실제 타입으로 변환 후 사용해야하는 경우
	// 인터페이스 타입 -> 형변환
	// interfaceVal.(type)

	var a interface{} = 15 // 빈 인터페이스의 최초의 타입이 선언된 타입이다
	b := a
	c := a.(int)
	// d := a.(float64)

	fmt.Println("Ex1 : ", a)
	fmt.Println("Ex1 : ", reflect.TypeOf(a))
	fmt.Println("Ex1 : ", b)
	fmt.Println("Ex1 : ", reflect.TypeOf(a))
	fmt.Println("Ex1 : ", c)
	fmt.Println("Ex1 : ", reflect.TypeOf(c))
	// fmt.Println("Ex1 : ", d)
	// fmt.Println("Ex1 : ", reflect.TypeOf(d))
	fmt.Println()

	// 예제2 - 저장된 실제 타입 검사
	if v, ok := a.(int); ok {
		fmt.Println("Ex2 : ", v, ok)
	}
}
