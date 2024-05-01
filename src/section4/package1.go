package main

import (
	"fmt"
	"os"
)

func main() {
	// 패키지 : 코드구조화(모듈화) 및 재사용
	// 패키지명 : 디렉토리 명
	// 같은 패키지 내 -> 소파일들은 디렉토리명을 패키지명으로 사용한다.
	// 네이밍 규칙 : 소문자 - private, 대문자 - public
	// package main는 특별하게 인식 -> 컴파일러는 프로그램의 시작점(entry point)로 인식

	var name string
	fmt.Println("이름은? : ")
	fmt.Scanf("%s", &name) // Scanf 사용자로부터 입력 받은 값
	fmt.Fprintf(os.Stdout, "HI %s\n", name)
}
