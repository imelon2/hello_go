// 고루틴 동기화 기초(4)

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 이론적인 뮤텍스 : 상호 배제
// Thread(고루틴) 들이 서로 running time에 서로 영향을 주지 않게
// 즉, 단독으로 실행되게 하는 기술

// RWMutex : 쓰기 Lock -> 쓰기 시도 중에는 다른 곳에서 이전 값을 읽으면 x
//
//	읽기 Lock -> 쓰기 Lock 전부 방어(방지)
//
// RMutex : 읽기 Lock -> 읽기 시도 중에 값 변경 방지
//
//	쓰기 Lock 방어(방지)
func main() {
	// 동기화 사용하는 경우 예제
	// 쓰기 읽기 동작 순서가 일정하지 않아, 잘못된 오류를 반환할 수있다.

	// 시스템 전체 cpu 사용
	runtime.GOMAXPROCS(runtime.NumCPU())
	data := 0 // 전역변수(공유데이터)
	mutex := new(sync.RWMutex)

	go func() {
		for i := 1; i <= 10; i++ {
			mutex.Lock() // 쓰기 잠금
			data += 1
			fmt.Println("Write : ", data)
			time.Sleep(200 * time.Microsecond)
			mutex.Unlock() // 쓰기 잠금 해제
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			// 읽기 뮤텍스 잠금
			mutex.RLock()
			fmt.Println("Read : ", data)
			time.Sleep(1 * time.Second)
			mutex.RUnlock()
			// 읽기 뮤텍스 해제
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			// 읽기 뮤텍스 잠금
			mutex.RLock()
			fmt.Println("Read2 : ", data)
			time.Sleep(1 * time.Second)
			mutex.RUnlock()
			// 읽기 뮤텍스 해제
		}
	}()

	time.Sleep(5 * time.Second)
}
