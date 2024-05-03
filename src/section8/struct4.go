// 구조체 기본(4)
package main

import "fmt"

func main() {
	// 구조체 익명 선언 및 활동

	// 예제1
	car1 := struct{ name, color string }{"520d", "red"}

	fmt.Println("Ex1 : ", car1)
	fmt.Printf("Ex1 : %#v \n", car1)

	// 예제2
	cars := []struct{ name, color string }{{"520d", "red"}, {"520d", "white"}, {"520d", "blue"}}

	for _, c := range cars {
		fmt.Printf("(%s, %s) ------- (%#v) \n", c.name, c.color, c)
	}
}
