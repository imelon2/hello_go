// 사용자 패키지 설치 및 활용 예제

package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

// 외부 패키지 설치
// (1) import 선언 후 폴더 이동 후 go get 설치
// (2) go get 패키기 주소 설치
func main() {
	// 선언 후 go get 패키지 설치
	xfile := "sample.xlsx"

	xlsx, err := xlsx.OpenFile(xfile)
	if err != nil {
		panic("Excel Loads Error")
	}

	for _, sheet := range xlsx.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				fmt.Printf("%s\t", cell)
			}
			fmt.Println()
		}
	}
}
