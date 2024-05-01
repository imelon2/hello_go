package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 문자열
	// 큰 다움표 "" '' , 백스쿼트 ``
	// char 없음 -> rune(=int32)로 문자코드 표현
	// 자주 사용하는 escape : \\, \' ,\"

	// 예제1
	// str1 == str2
	var str1 string = "$HOME\\Documents\\me\\hello_go"
	str2 := `$HOME\Documents\me\hello_go`

	fmt.Println("ex1 : ", str1)
	fmt.Println("ex1 : ", str2)

	// 예제2
	// len()은 총 byte 수를 길이로 취급한다.
	// en: 1 byte | kn: 3 byte
	var str3 string = "안녕하세요"
	var str4 string = "hello world"

	fmt.Println("ex2 : ", len(str3))
	fmt.Println("ex2 : ", len(str4))

	// 예제3
	// 문자열의 길이(length)는 utf8.RuneCountInString()로 구할 수 있다
	// 문자열의 길이(length)는 len([]rune())로 구할 수 있다.
	fmt.Println("ex3 : ", utf8.RuneCountInString(str3))
	fmt.Println("ex3 : ", utf8.RuneCountInString(str4))
	fmt.Println("ex3 : ", len([]rune(str3)))
}
