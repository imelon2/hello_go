package main

import "fmt"

func main() {
	// 맵(map)
	// 맵 : 해시테이블, 딕셔너리,Key-Value 자료저장
	// Reference Type
	// 비교연산자 사용x
	// 특징 : 참조타입 key로 사용 불가, 값(Value)으로 모든 타입 사용 가능
	// make() 및 축양(리터럴)로 초기화 가능
	// 순서 없음!!

	// 예제1 - 선언
	// `map[(Key type)] (Value type)`
	var map1 map[string]int = make(map[string]int) // 정석
	var map2 = make(map[string]int)                // 타입 자동 할당

	map3 := make(map[string]int)

	fmt.Println("ex1 : ", map1)
	fmt.Println("ex1 : ", map2)
	fmt.Println("ex1 : ", map3)
	fmt.Println()

	// 예제2 - JSON
	map4 := map[string]int{} // Json 형태
	map4["apple"] = 25
	map4["banana"] = 40
	map4["orange"] = 33

	map5 := map[string]int{
		"apple":  15,
		"banana": 40,
		"orange": 23,
	}

	map6 := make(map[string]int, 10)
	map6["apple"] = 25
	map6["banana"] = 40
	map6["orange"] = 33

	// 여러번 실행해보면 순서x 확인 가능
	fmt.Println("ex2 : ", map4)
	fmt.Println("ex2 : ", map5)
	fmt.Println("ex2 : ", map6)
	fmt.Println("ex2 : ", map6["orange"])
	fmt.Println("ex2 : ", map6["apple"])
	fmt.Println("ex2 : ", map6["invaild"])
}
