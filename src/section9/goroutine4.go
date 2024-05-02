// Goroutine 기초 (4)
package main

import (
	"fmt"
	"runtime" // 운영체제 명령어
	"time"
)

func main() {
	// 고루틴
	// 클로저 사용 예제

	// 예제1
	runtime.GOMAXPROCS(4)
	s := "GoRoutine Closure : "
	for i := 0; i < 1000; i++ {
		// 반복문 클로저는 일반적으로 즉시 실행
		// BUT 고루틴 클로저는 가장 나중에 실행(반복문 종료 후)
		// 반복문이 끝남과 동시에 한번에 실행(다중 쓰레드)
		go func(n int) {
			fmt.Println(s, n, " - ", time.Now())
		}(i)
	}

	time.Sleep(3 * time.Second)
}
