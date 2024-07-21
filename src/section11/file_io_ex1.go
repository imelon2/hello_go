// 파일 I/O (1)

package main

import (
	"fmt"
	"os"
)

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}

// 파일 읽기 쓰기 -> ioutil 패키지 활용
// 더욱 편리하고 직관적으로 파일 읽고 쓰기
// WriteFile(), ReadFile(), ReadAll()
func main() {
	s := "Hello Golang\nFile Write Test! By ioutils\n"

	// 파일모드(chmod, chown, chgrp) -> permission
	// 읽기 : 4 , 쓰기 : 2, 실행 : 1
	// 각 index 마다 숫자 더하기
	// 소유자, 그룹, 기타 사용자 순서
	// example
	// (777) 소유자, 그룹, 기타 사용자 모드 권한 부여
	// 644 : 소유자(읽기,쓰기) 그룹, 기타 사용자 모드(읽기)
	// https://pkg.go.dev/os@go1.22.5#FileMode

	err := os.WriteFile("test_write1.txt", []byte(s), os.FileMode(0644))
	errCheck(err)

	data, err := os.ReadFile("sample.txt")
	errCheck(err)

	fmt.Println("=================================")
	fmt.Println(string(data))
	fmt.Println("=================================")

}
