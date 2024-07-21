// 고루틴 동기화 기초(3)

package main

import (
	"fmt"
	"runtime"
	"time"
)

// 이론적인 뮤텍스 : 상호 배제
// Thread(고루틴) 들이 서로 running time에 서로 영향을 주지 않게
// 즉, 단독으로 실행되게 하는 기술
func main() {
	// 동기화 사용하지 않은 경우 예제
	// 쓰기 읽기 동작 순서가 일정하지 않아, 잘못된 오류를 반환할 수있다.

	// 시스템 전체 cpu 사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	data := 0 // 전역변수(공유데이터)

	go func() {
		for i := 1; i <= 10; i++ {
			data += 1
			fmt.Println("Write : ", data)
			time.Sleep(200 * time.Microsecond)
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println("Read : ", data)
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println("Read2 : ", data)
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(5 * time.Second)
}
