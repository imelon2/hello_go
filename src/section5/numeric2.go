package main

import (
	"fmt"
)

func main() {
	// 정수형 문자
	// 예제1 - 아스키(영문)
	var char1 uint8 = 72
	var char2 byte = 0110
	var char3 byte = 0x48

	// 예제2 - 유니코드(한글)
	var char4 rune = 50556
	var char5 rune = 0142574 // 44032
	var char6 rune = 0xC57C  //44032

	// string로 출력
	fmt.Printf("%c %c %c\n", char1, char2, char3)
	// 숫자로 표현해
	fmt.Printf("%d %d %d\n", char1, char2, char3)
	// 숫자, 8진수, 16진수로 표현해
	fmt.Printf("%d %o %x\n", char1, char2, char3)

	fmt.Printf("%c %c %c\n", char4, char5, char6)
	fmt.Printf("%d %d %d\n", char4, char5, char6)
	fmt.Printf("%d %o %x\n", char4, char5, char6)
}
