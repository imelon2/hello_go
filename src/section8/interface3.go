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
func (d Dog) bite() {
	fmt.Println(d.name, " : Dog bites!")
}

func (d Dog) sound() {
	fmt.Println(d.name, " : 멍멍!!")
}
func (d Dog) run() {
	fmt.Println(d.name, " : 뛰어!!!")
}

// 구조체 Cat 메소드 구현
func (c Cat) bite() {
	fmt.Println(c.name, " : 깨물기!")
}

func (c Cat) sound() {
	fmt.Println(c.name, " : 냥냥!!")
}

func (c Cat) run() {
	fmt.Println(c.name, " : 폴짝!!!")
}

// 동물의 행동 인터페이스 선언
type Behavior1 interface {
	bite()
	sound()
	run()
}

// 인터페이스를 파라미터로 받는다.
func act(animal Behavior1) {
	animal.bite()
	animal.sound()
	animal.run()
}

func main() {
	dog := Dog{"강아지", 10}
	cat := Cat{"고양이", 10}

	act(dog)

	act(cat)
}
