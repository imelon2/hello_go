// 채널(Channel) 기초(4)
package main

import (
	"fmt"
	"runtime"
)

// 비동기 버퍼 사용
// 버퍼 : 발신 -> 가득차면 대기, 비어있으면 작동
//
//	: 수신 -> 비어있으면 대기, 가득차면 작동
func main() {

	runtime.GOMAXPROCS(1)

	ch := make(chan bool, 2)
	cnt := 12

	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true
			fmt.Println("Go : ", i)
		}
	}()

	for i := 0; i < cnt; i++ {
		<-ch
		fmt.Println("Main : ", i)
	}
}
