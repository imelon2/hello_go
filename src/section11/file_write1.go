// 파일 쓰기 (1)
package main

import (
	"fmt"
	"os"
)

// 에러 체크 방식(1)
func errCheck1(e error) {
	if e != nil {
		panic(e)
	}
}

// 에러 체크 방식(2)
func errCheck2(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Create : 새 파일 작성 및 파일 열기
	// Close : 리소스 닫기
	// Write, WriteString, WriteAt : 파일 쓰기
	// 각 운영체제 권한 주의(오류 메시지 확인)
	// 예외 처리 정말 중요!!

	// 파일 쓰기 예제
	file, err := os.Create("test_write.txt")
	errCheck1(err)

	// 리소스 해제!!
	defer file.Close()

	// 쓰기 예제 (1)
	s1 := []byte{115, 111, 109, 120, 112}

	// return size
	//  문자열 -> byte 슬라이스 형으로 변환 후 쓰기
	n1, err := file.Write(s1)
	// n1, err := file.Write([]byte(s1))
	errCheck2(err)

	fmt.Printf("쓰기 작업 완료(1) (%d bytes) \n", n1)

	// Write Commit 작업(stable)
	file.Sync()

	// 쓰기 예제 (2)
	s2 := "Hello Golang!!! \n File Write Test!! - 2\n"
	n2, err := file.WriteString(s2)
	errCheck2(err)

	fmt.Printf("쓰기 작업 완료(2) (%d bytes) \n", n2)
	file.Sync()

	// 쓰기 예제 (3)
	s3 := "Test WriteAt! - 3\n"
	n3, err := file.WriteAt([]byte(s3), 10) // space 70번 위치한 곳에서 작성
	errCheck2(err)

	fmt.Printf("쓰기 작업 완료(3) (%d bytes) \n", n3)
	file.Sync()

	// 쓰기 예제 (4)
	s4 := "Hello Golang ! \n File Write Test - 4\n"
	n4, err := file.WriteString(s4)
	errCheck2(err)

	fmt.Printf("쓰기 작업 완료(4) (%d bytes) \n", n4)
	file.Sync()
}
