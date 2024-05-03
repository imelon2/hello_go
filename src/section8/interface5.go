package main

import (
	"fmt"
)

type Dog struct {
	name   string
	weight int
}

type Cat struct {
	name   string
	weight int
}

// 어떠한 타입이라도 유연하게 매개변수를 받을 수 있다
// type any처럼 사용 가능
func printValue(s interface{}) {
	fmt.Println("Ex1 : ", s)
}

func main() {
	// 인터페이스 활용(빈 인터페이스)
	// 함수내에서 어떠한 타입이라도 유연하게 매개변수를 받을 수 있다(만능) -> 모든 타입 지정 가능

	dog := Dog{"강아지", 10}
	cat := Cat{"고양이", 10}

	printValue(dog)
	printValue(cat)

	printValue(15)
	printValue("Animal")
	printValue([5]Dog{})
}
