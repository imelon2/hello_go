package main

import (
	"fmt"
)

// init : package 로드 시, 가장 먼저 호출
// 여러개 사용가능하며, 위에서 부터 순서대로 실행됨 -> 단 여러개 쓰는건 보기가 안좋음
func init() {
	fmt.Println("Init Method Start!!")
}

func main() {
	fmt.Println("Main Method Start!!")
}
