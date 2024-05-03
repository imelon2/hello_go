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

// 구조체 Dog 메소드 구현
func (d Dog) run() {
	fmt.Println(d.name, " : 뛰어!!!")
}

// 구조체 Cat 메소드 구현
func (c Cat) run() {
	fmt.Println(c.name, " : 폴짝!!!")
}

func act(animal interface{ run() }) {
	animal.run()
}

func main() {
	// 익명 인터페이스 사용 예지(즉시 사용)

	dog := Dog{"강아지", 10}
	cat := Cat{"고양이", 10}

	act(dog)

	act(cat)
}
