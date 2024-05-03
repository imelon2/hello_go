package main

import (
	"fmt"
)

type test interface {
}

func main() {

	// 인터페이스
	// 객체의 동작을 표현 - 골격
	// 단순히 동작에 대한 방법만 표시
	// 추상화 제공
	// 인터페이스의 메소드를 구현한 타입은 인터페이스로 사용 가능
	// 유연하게 사용
	// [중요] 덕타이핑 : Go언어의 독창적인 특성

	// 예제1
	/*
		type 인터페이스명 interface {
			메서드1()
			메서드2()
		}
	*/

	var t test
	fmt.Println("Ex1 : ", t)
}
