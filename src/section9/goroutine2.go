// Goroutine 기초 (2)
package main

import (
	"fmt"
	"time"
)

func exe(name string) {
	fmt.Println(name, " start : ", time.Now())
	for i := 0; i < 100; i++ {
		fmt.Println(name, " >>> ", i)
	}
	fmt.Println(name, " end : ", time.Now())
}

func main() {
	// 고루틴
	exe("t1")

	// 예제1
	fmt.Println("Main Routine Start : ", time.Now())
	// 어떤 쓰레드가 먼저 실행되는지, 순서가 뒤죽박죽
	go exe("t2")
	go exe("t3")

	time.Sleep(4 * time.Second)
	fmt.Println("Main Routine End : ", time.Now())
}
