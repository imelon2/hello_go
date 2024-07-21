// 고루틴 동기화 기초(2)

package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 구조체 선언(공유데이터)
type count1 struct {
	num   int
	mutex sync.Mutex
}

func (c *count1) increment1() {
	c.mutex.Lock()
	c.num += 1
	c.mutex.Unlock()
}

func (c *count1) result1() {
	fmt.Println(c.num)
}

// 고루틴 동기화
// 실행 흐름 제어 및 변수 동기화 가능
// 공유 데이터 보호가 가장 중요!!
// 만약, 멀티 스레드 실행 시 데이터 또는 변수를 공유할 경우, 경쟁 상태가 될 수 있음
// 뮤텍스(Mutex) : 경쟁 상태를 방지하는 기능
// sync.Mutex 선언 후, Lock & Unlock 사용
func main() {

	// 동기화 사용한 경우 예제
	// 시스템 전체 cpu 사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := count1{num: 0}
	done := make(chan bool)

	for i := 1; i <= 10000; i++ {
		go func() {
			c.increment1()
			done <- true
			runtime.Gosched() // 다른 Cpu 양보
		}()
	}

	for i := 1; i <= 10000; i++ {
		<-done
	}

	c.result1()
}
