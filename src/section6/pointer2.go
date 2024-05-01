package main

import "fmt"

func main() {
	// 예제1
	i := 7  // i : int type
	p := &i // p : pointer type

	fmt.Println("Ex1 : ", i, &i)
	fmt.Println("Ex1 : ", p, *p)
	fmt.Println()

	*p++
	fmt.Println("Ex1 : ", i, &i)
	fmt.Println("Ex1 : ", p, *p)
	fmt.Println()

	*p = 77777
	fmt.Println("Ex1 : ", i, &i)
	fmt.Println("Ex1 : ", p, *p)
	fmt.Println()

	i = 99999
	fmt.Println("Ex1 : ", i, &i)
	fmt.Println("Ex1 : ", p, *p)
	fmt.Println()
}
