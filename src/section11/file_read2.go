// 파일 읽기 (2)
package main

import (
	"bufio"
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
	// CSV 파일 읽기 예제
	// 파일 용량이 큰 경우 `bufio` 사용 권장

	// 파일 생성
	file, err := os.Open("sample.csv")
	errCheck1(err)

	// 리소스 해제
	defer file.Close()

	// CSV Reader 생성
	// bufio 사용 권장
	rr := csv.NewReader(bufio.NewReader(file))

	// CSV 내용 읽기 (1)
	row1, err1 := rr.Read() //  1개의 Row 단위로 읽기
	errCheck2(err1)
	// (중요) 읽을 때 마다 위치 이동!!
	row2, err2 := rr.Read() //  1개의 Row 단위로 읽기
	errCheck2(err2)

	fmt.Println(" === CSV Read Example ===")
	// fmt.Println(row)
	fmt.Println(row1[1], " ", row1[1:5])
	fmt.Println(row2[1], " ", row2[1:5])
	fmt.Println("==========================================")

	// CSV 내용 읽기 (2)
	rows, err := rr.ReadAll() // 전체 Row 읽기
	errCheck2(err)

	fmt.Println(" === CSV Read Example ===")
	fmt.Println(rows[5][1])
	// fmt.Println(rows)
	fmt.Println("==========================================")

}
