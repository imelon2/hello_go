// Go Error 처리 기초 (1)

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 애러처리 : 소프트웨어의 품질 향상 가장 중요한것!!
	// GO에서는 기본적으로 error 패키지에서 error 인터페이스 제공
	// type error interface { Error() starint }
	// 즉, Error 메서드를 구현하면 사용자 정의 애러 처리 제작이 기능하다!

	// 기본적으로 메서드마다 return Type은 2개(return data, err)
	// 주로 Errorf(에러 타입 리턴) 메소드, Fatal(프로그램 종료) 메소드를 통해 애러 출력

	// 기본적인 메서드 에러 처리 예제
	f, err := os.Open("unnamefile")
	if err != nil {
		// log.Fatal(err.Error()) // 방법 1
		log.Fatal(err) // 방법 2
	}

	fmt.Println("===============")
	fmt.Println("ex1 : ", f.Name())
}
