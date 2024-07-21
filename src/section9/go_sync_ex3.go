// 고루틴 동기화 고급(3)

package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 원자성(Atomic) 사용 -> 기능적으로 분할 불가능한 완전 보증된 일련의 조작
// 모든 조작이 완료될 때 까지 다른 프로세스 개입불가
// ex. 트랜잭션
// sync/atomic에서 완자적 연산자 제공
// https://pkg.go.dev/sync/atomic
// 주로 공용 변수에 관한 계산 사용

func main() {
	// 원자성 사용 안할 경우 예제
	// 시스템 전체 cpu 사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 공융 변수
	var cnt int64 = 0
	wg := new(sync.WaitGroup)

	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func(n int) {
			cnt++
			wg.Done()
		}(i)
	}

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func(n int) {
			cnt -= 1
			wg.Done()
		}(i)
	}

	// Add(7000) == Done(7000) 횟수가 같을 때 까지 대기
	wg.Wait()
	fmt.Println("cnt : ", cnt)
	fmt.Println("WaitGroup End!")
}
