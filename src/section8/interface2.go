package main

import (
	"fmt"
)

type Dog struct {
	name   string
	weight int
}

func (d Dog) bite() {
	fmt.Println(d.name, " bites !!")
}

// 동물의 행동 인터페이스 선언
type Behavior interface {
	bite()
}

func main() {
	// 인터페이스 구현 예제
	// 예제1

	dog1 := Dog{"poll", 10}

	var inter1 Behavior
	inter1 = dog1
	inter1.bite()

	// 예제2
	dog2 := Dog{"marry", 12}
	inter2 := Behavior(dog2)
	inter2.bite()

	// 예제3
	inters := []Behavior{dog1, dog2}
	for i, _ := range inters {
		inters[i].bite()
	}

	for _, val := range inters {
		val.bite()
	}

	// 예제4

}
