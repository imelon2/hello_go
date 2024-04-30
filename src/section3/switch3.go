package main

import (
	"fmt"
)

func main() {
	// example 1
	a := 30 / 15
	switch a {
	case 2, 4, 6:
		fmt.Println("a -> ", a, "는 짝수")
	case 1, 3, 5:
		fmt.Println("a -> ", a, "는 홀수")
	}

	// example 2 : fallthrough
	switch e := "python"; e {
	case "java":
		fmt.Println("Java!!")
		fallthrough
	case "go":
		fmt.Println("go")
		// fallthrough
	case "python":
		fmt.Println("python")
		fallthrough
	case "ruby":
		fmt.Println("ruby")
		// fallthrough
	case "c":
		fmt.Println("c")
	}
}
