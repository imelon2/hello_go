package main

import "fmt"

// FILO
func stack() {
	for i := 1; i <= 10; i++ {
		defer fmt.Println("Ex1 : ", i)
	}
}

func main() {
	stack()
}
