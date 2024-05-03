package main

import "fmt"

func ex_f1() {
	fmt.Println("f1 : start!!")
	defer ex_f2()
	fmt.Println("f1 : end!!")
}

func ex_f2() {
	fmt.Println("f2 : called!!")
}

func main() {
	// Defer - 지연 실행
	// Defer를 호출한 함수가 종료되기 직전에(마지막에) 호출된다.
	// Finally와 비슷
	// 주로 리소스 반환, 열린 파일 닫기, Mutex 잠금 해제

	// 예제1
	ex_f1()
}
