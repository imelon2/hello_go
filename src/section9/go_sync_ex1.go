// 고루틴 동기화 고급(1)

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func onceTest() {
	fmt.Println("Once Test Excute!!")
}

// Once : 한번만 실행(주로 초기화 사용)
// Do : 실행
func main() {
	// 시스템 전체 cpu 사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	once := new(sync.Once)

	for i := 0; i < 5; i++ {
		go func(n int) {
			fmt.Println("Gorutine : ", n)
			once.Do(onceTest)
		}(i)
	}

	time.Sleep(2 * time.Second)
}
