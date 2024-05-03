// 구조체 기본(6)
package main

import (
	"fmt"
)

type Car1 struct {
	name    string "차량명"
	color   string "색상"
	company string "제조사"
	detail  spec   "상세"
}

type spec struct {
	length int "전장"
	height int "전고"
	width  int "전축"
}

func main() {
	// 중첩 구조체

	// 예제1
	car1 := Car1{
		"520d",
		"siler",
		"bmw",
		spec{400, 100, 200},
	}

	fmt.Println("Ex1 : ", car1.name)
	fmt.Println("Ex1 : ", car1.color)
	fmt.Println("Ex1 : ", car1.company)
	fmt.Printf("Ex1 : %#v \n", car1.detail)

	// 예제2
	// 내부 spec 구조체 필드 값 출력
	fmt.Println("Ex2 : ", car1.detail.height)
	fmt.Println("Ex2 : ", car1.detail.length)
	fmt.Println("Ex2 : ", car1.detail.width)
}
