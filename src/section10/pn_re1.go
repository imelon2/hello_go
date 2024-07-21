// Panic & Recover (1)
package main

import (
	"fmt"
	"log"
)

func main() {
	// Go panic 함수
	// 사용자가 에러 발생 가능
	// Panic : 호출 즉시, 해당 메소드를 즉시 중지시키고 defer 함수를 호출하고 자기자신을 호출한 곳으로 리턴
	// 런타임 이외에 사용자가 코드 흐름에 따라 에러를 발생 시킬때 중요!!
	// ㄴ 문법적인 에러는 아니지만, 논리적인 코드 흐름에 따른 에러 발생 처리 기능
	// 즉, Runtime 과정에서 발생한 애러!!!
	// 복구까지 가능
	fmt.Println("Start Main")
	// 방법 (1)
	// panic("Error occurred : User Stoped!")
	// 방법 (2)
	log.Panic("Error occurred : User Stoped!")

	fmt.Println("End Main")
}
