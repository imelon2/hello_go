// Goroutine 기초 (1)
package main

import (
	"fmt"
	"time"
)

func exe1() {
	fmt.Println("exe1 func start", time.Now())
	time.Sleep(1 * time.Second) // 1초
	fmt.Println("exe1 func end", time.Now())
}

func exe2() {
	fmt.Println("exe2 func start", time.Now())
	time.Sleep(1 * time.Second) // 1초
	fmt.Println("exe2 func end", time.Now())
}

func exe3() {
	fmt.Println("exe3 func start", time.Now())
	time.Sleep(1 * time.Second) // 1초
	fmt.Println("exe3 func end", time.Now())
}

func main() {
	// 고루틴
	// 타 언어의 쓰레드(Tread)와 비슷한 기능
	// 생성 방법 매우 간단, 리소스 매우 적게 사용
	// 즉, 수많은 고루틴 동시 생성 및 실행 가능
	// 비동기적 함수 루틴 실행 -> 채널(?)을 통한 통신 가능
	// 공유 메모리 사용 시에 정확한 동기화 코딩 필요

	// 싱글 루틴에 비해 항상 빠른 처리는 아니다!!

	// 멀티 쓰레드 장단점
	// 장점 : 응답성 향상, 효율적인 자원 공유, 작업 분리된 간결한 코드
	// 단점 : 구현하기 어려움, 테스트 및 디버깅 어려움, 전체 프로세스의 사이트 이펙트 발생, 성능 저하, 동기화 코딩 필수
	// 		 데드락...

	// 데몬스레드 : 메인 쓰레드가 끝나면 다른 쓰레드도 종료
	exe1()

	fmt.Println("Main Routine Start", time.Now())
	// func Main이 1초만에 끝나버리면 다른 고루틴은 같이 종료된다.
	go exe2() // demin tread
	go exe3() // demin tread

	time.Sleep(4 * time.Second)
	fmt.Println("Main Routine Enf", time.Now())
}
