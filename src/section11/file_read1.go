// 파일 읽기 (1)
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

// Open : 기존의 파일 열기
// Close : 리소스 닫기
// Read, ReadAt : 파일 읽기
// 각 운영 체제 권한 주의!!

// 탐색 Seek 중요!!
func main() {
	file, err := os.Open("sample.txt")
	errCheck1(err)

	defer file.Close()

	// 파일 열기 예제 (1)
	fileInfo, err := file.Stat() // 파일 사이즈 확인
	errCheck2(err)

	// 내용(사이즈) 만큼 Slice 생성
	fd1 := make([]byte, fileInfo.Size())

	ct1, err := file.Read(fd1)

	fmt.Println("파일 정보 출력 (1) : ", fileInfo, "\n")
	fmt.Println("파일 이름 (2) : ", fileInfo.Name(), "\n")
	fmt.Println("파일 크기 (3) : ", fileInfo.Size(), "\n")
	fmt.Println("파일 최종 수정 시간 (4) : ", fileInfo.ModTime(), "\n")
	fmt.Printf("읽기 작업 (1) 완료 (%d bytes)\n\n", ct1)
	fmt.Println(string(fd1))

	fmt.Println("================================")

	// 파일 열기 예제 (2) 탐색 : Seek(offset)
	// offset부터 읽은 후, 어디서부터 읽을지 위치 지정 파라미터
	// 0: 처음 위치
	// 1: 현재 위치
	// 2: 마지막 위치
	o1, err := file.Seek(20, 0)
	errCheck1(err)

	fd2 := make([]byte, 20)
	ct2, err := file.Read(fd2)
	errCheck2(err)

	fmt.Printf("읽기 작업 (2) 완료 (%d bytes) (%d offset) \n\n", ct2, o1)
	fmt.Println(string(fd2))

	fmt.Println("================================")

	// 파일 열기 예제 (2) 탐색
	// offset부터 읽은 후, 어디서부터 읽을지 위치 지정 파라미터
	// 0: 처음 위치
	// 1: 현재 위치
	// 2: 마지막 위치
	o2, err := file.Seek(0 /* offset */, 0)
	errCheck1(err)

	fd3 := make([]byte, 50)
	ct3, err := file.ReadAt(fd3, 8 /* offset */)
	errCheck2(err)

	fmt.Printf("읽기 작업 (3) 완료 (%d bytes) (%d offset) \n\n", ct3, o2)
	fmt.Println(string(fd3))
	fmt.Println("================================")
}
