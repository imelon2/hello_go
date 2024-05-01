package main

import (
	"fmt"
)

func main() {
	var n1 int = 12
	var n2 float32 = 8.2

	// 두개 값이 다름
	fmt.Println("ex3 : ", float32(n1)+n2)
	fmt.Println("ex3 : ", n1+int(n2))
}
