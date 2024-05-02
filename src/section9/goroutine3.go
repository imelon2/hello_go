// Goroutine 기초 (3)
package main

import (
	"fmt"
	"math/rand"
	"runtime" // 운영체제 명령어
	"time"
)

func exe123(name int) {
	r := rand.Intn(100) // 100 미만 랜덤
	fmt.Println(name, " start : ", time.Now())
	for i := 0; i < 100; i++ {
		fmt.Println(name, " >>>> ", r, i)
	}
	fmt.Println(name, " end : ", time.Now())

}

func main() {
	// 고루틴
	// 멀티 코어(다중 cpu) `최대한` 활용

	currCPU := runtime.NumCPU() // 내 PC CPU 개수
	runtime.GOMAXPROCS(currCPU)
	fmt.Println("Current System CPU : ", runtime.GOMAXPROCS(0)) // 설정값 출력

	// 예제1
	fmt.Println("Main Routine Start : ", time.Now())
	for i := 0; i < 100; i++ {
		go exe123(i) // 고루틴 100 개
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Main Routine End : ", time.Now())
}
