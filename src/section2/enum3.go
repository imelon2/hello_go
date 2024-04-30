package main

import "fmt"

func main() {
	const (
		_ = iota
		A // 1
		_
		C // 3
	)

	fmt.Println(A, C)
}
