package main

import "fmt"

func main() {
	const (
		A = iota // 0
		B        // 1
		C        // 2
	)

	fmt.Println(A, B, C)

	const (
		D = iota * 10 // 0
		E             // 10
		F             // 20
	)

	fmt.Println(D, E, F)
}
