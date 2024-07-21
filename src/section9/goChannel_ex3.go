// 채널(Channel) 심화(3)
package main

import (
	"fmt"
)

func receiveOnly1(cnt int) <-chan int {
	sum := 1
	tot := make(chan int)
	go func() {
		for i := 1; i <= cnt; i++ {
			sum += i
		}
		tot <- sum
		tot <- 777
		tot <- 777
		close(tot)
	}()
	return tot
}
func total(c <-chan int) <-chan int {
	tot := make(chan int)
	go func() {
		a := 0
		for i := range c {
			a = a + i
		}
		tot <- a
	}()

	return tot
}

func main() {
	c := receiveOnly1(100) // 채널 반환
	output := total(c)     // 채널 전달후 반환

	fmt.Println("ex1 : ", <-output)
}
