// 채널(Channel) 기초(2)
package main

import (
	"fmt"
)

func rangeSum(rg int, c chan int) {
	sum := 0
	for i := 1; i <= rg; i++ {
		sum += i
	}

	c <- sum
}

func main() {
	c := make(chan int)

	go rangeSum(1000, c)
	go rangeSum(7000, c)
	go rangeSum(5000, c)

	result1 := <-c
	result2 := <-c
	result3 := <-c

	fmt.Println("ex1 : ", result1)
	fmt.Println("ex1 : ", result2)
	fmt.Println("ex1 : ", result3)
}
