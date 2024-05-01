package main

import (
	"fmt"
	"math"
)

// 숫자 연산 - 산술, 비교
// 타입이 같아야만 가능(엄격)
// 다른 타입은 `반드시` 형 변환 후 연산
func main() {
	// 예제1
	var n1 uint8 = math.MaxInt8
	var n2 uint16 = math.MaxInt16
	var n3 uint32 = math.MaxInt32
	var n4 uint64 = math.MaxInt64

	fmt.Println("n1 : ", n1)
	fmt.Println("n2 : ", n2)
	fmt.Println("n3 : ", n3)
	fmt.Println("n4 : ", n4)

	fmt.Println("float32 : ", math.MaxFloat32)
	fmt.Println("float64 : ", math.MaxFloat64)

	n5 := 100000 // int
	n6 := int16(1000)
	n7 := uint8(30)

	// fmt.Println(n5 + n6) // type error
	fmt.Println("cul1 : ", n5+int(n6))
	// fmt.Println("cul2 : ",n6 + n7) // type error
	fmt.Println("cul2 : ", n6+int16(n7))

}
