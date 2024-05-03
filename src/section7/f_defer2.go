package main

import "fmt"

func sayHello(msg string) {
	defer func() {
		fmt.Println(msg)
	}()

	func() {
		fmt.Println("HI!!!")
	}()
}

func main() {
	// 예제1
	sayHello("Golang!!")

}
