// 고루틴 동기화 고급(2)

package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 대기 그룹
// 실행 흐름 변경(고루틴이 종료 될 때 까지 대기 기능)
// WaitGroup : Add(고루틴 추가)
// Done(작업 종료 알림)
// Wait (고루틴 종료 까지 대기)
// Add로 추가된 고루틴 개수와 Done으로 종료되는 알림 횟수와 같아야 한다. (Add == Done)
func main() {
	// 시스템 전체 cpu 사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg := new(sync.WaitGroup)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			fmt.Println("WaitGroup : ", n)
			wg.Done()
		}(i)
	}

	// Add == Done 횟수가 같을 때 까지 대기
	wg.Wait()
}
