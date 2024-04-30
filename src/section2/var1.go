package main

import "fmt"

func main() {
	// 초기화
	// 정수타입, 실수(소수점), 문자열, Boolean

	var a int
	var b string
	var c, d, e int
	var f, g, h int = 1, 2, 3
	var i float32 = 11.4
	var j string = "Hi! Golang!!"
	var k = 4.74 // 자동 선언
	var l = "HI Seoul!!"
	var m = true

	a = 4
	b = "Hello Go"

	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
	fmt.Println("c : ", c)
	fmt.Println("d : ", d)
	fmt.Println("e : ", e)
	fmt.Println("f : ", f)
	fmt.Println("g : ", g)
	fmt.Println("h : ", h)
	fmt.Println("i : ", i)
	fmt.Println("j : ", j)
	fmt.Println("k : ", k)
	fmt.Println("l : ", l)
	fmt.Println("m : ", m)

	// a :  0
	// b :
	// c :  0
	// d :  0
	// e :  0
	// f :  1
	// g :  2
	// h :  3
	// i :  11.4
	// j :  Hi! Golang!!
	// k :  4.74
	// l :  HI Seoul!!
	// m :  true
}
