package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("ex1 : ", i)
	}

	loc := []string{"s", "b", "i"}
	for index, name := range loc {
		fmt.Println("ex3 : ", index, name)
	}
	for _, name := range loc {
		fmt.Println("ex3 : ", name)
	}
}
