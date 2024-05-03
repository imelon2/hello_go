// 사용자 정의 타입(1)
package main

import "fmt"

type Car struct {
	name  string
	color string
	price rune
	tax   rune
}

// 일반 메소드
func Price(c Car) int {
	return int(c.price) + int(c.tax)
}

// 구조체 <-> 메소드 바인딩
func (c Car) Price() int {
	return int(c.price) + int(c.tax)
}

func main() {
	// Go는 객체 지향 타입을 `구조체`로 정의한다. (클래스, 상속 개념 없음)

	// 상태, 메소드 분리해서 정의(결합성x)
	// 사용자 정의 타입 이란? 구조체, 인터페이스, 기본타입, 함수 정의
	// 구조체와 메소드 연결을 통해 객체 지향 처럼 사용 가능

	// 예제1
	bmw := Car{
		name:  "520d",
		price: 500000000,
		color: "white",
		tax:   5000,
	}
	benz := Car{
		name:  "220d",
		price: 1000000000,
		color: "white",
		tax:   3000,
	}

	fmt.Println("Ex1 : ", bmw, &bmw)
	fmt.Println("Ex1 : ", benz, &benz)
	// fmt.Printf("Ex1 : %p", &bmw, "\n")
	// fmt.Printf("Ex1 : %p", &benz, "\n")

	// 예제2 - 일반 메소드
	fmt.Println("Ex2 : ", Price(bmw))
	fmt.Println("Ex2 : ", Price(benz))

	// 구조체 <-> 메소드 바인딩
	fmt.Println("Ex2 : ", bmw.Price())
	fmt.Println("Ex2 : ", benz.Price())

	// 예제3 - 서로 다른 메모리
	fmt.Print("Ex3 : ", &bmw == &benz) // false
}
