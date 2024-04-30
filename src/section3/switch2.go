package main

import (
	"fmt"
	"math/rand" // 랜덤 숫자 생성 함수
	"time"      //  시간 관련 함수
)

func main() {
	// rand.Intn(100) : 0 ~ 100 사이 랜덤 숫자

	rand.Seed(time.Now().UnixNano()) // 시드를 활용한 랜덤 숫자

	switch i := rand.Intn(100); {
	case i >= 50 && i < 100:
		fmt.Println("i -> ", i, "50 이상 100 미만")
	case i >= 25 && i < 50:
		fmt.Println("i -> ", i, "25 이상 50 미만")
	default:
		fmt.Println("없음")
	}
}
