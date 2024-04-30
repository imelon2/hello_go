package main

import "fmt"

func main() {
	var a int = 50
	b := 70

	// example 1
	if a >= 65 {
		fmt.Println("65 이상")
	} else {
		fmt.Println("65 미만")
	}

	// example 2
	if b >= 70 {
		fmt.Println("70 이상")
	} else {
		fmt.Println("70 미만")
	}
}
