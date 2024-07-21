// 파일 쓰기 (1)
package main

import (
	"encoding/csv"
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
	// CSV 파일 쓰기 예제
	// 패키지 저장소를 통해서 Excel 등 다양한 파일 형식 쓰기 읽기 가능
	// 패키지 Github 주소 : https://github.com/tealeg/xlsx
	// bufio : 파일 용량이 큰 경우 버퍼 사용 권장

	// 파일 생성
	file, err := os.Create("test_write.csv")
	errCheck1(err)

	// 리소스 해제
	defer file.Close()

	// csv writer 생성
	wr := csv.NewWriter(file)

	// 큰 파일을 사용할때
	// wr := csv.NewWriter(bufio.NewWriter(file))

	// csv 내용 쓰기
	wr.Write([]string{"Kim", "4.1"})
	wr.Write([]string{"Lee", "4.2"})
	wr.Write([]string{"Park", "4.3"})
	wr.Write([]string{"Jo", "4.4"})
	wr.Write([]string{"Hong", "4.5"})
	wr.Write([]string{"Choi", "4.6"})

	// 버퍼 -> 파일로 쓰기
	wr.Flush() // file.Close()랑 같음

	fi, err := file.Stat()
	errCheck1(err)

	fmt.Printf("CSV 쓰기 작업 후 파일 크기 (%d byte) \n", fi.Size())
	fmt.Println("CSV 파일명 : ", fi.Name())
	fmt.Println("운영체제 파일 권한: ", fi.Mode())
}
