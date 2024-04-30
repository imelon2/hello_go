package main

import "fmt"

func main() {
	// 짧은 선언
	// 정해진 함수 안에서만 사용 가능하며, 함수가 종료되면 메모르 초기회(전역변수 x)
	shortVar1 := 3
	shortVar2 := "Test"
	shortVar3 := false
	fmt.Println("short Var1 : ", shortVar1)
	fmt.Println("short Var2 : ", shortVar2)
	fmt.Println("short Var3 : ", shortVar3)

	// 선언 후 재할당 불가
	// shortVar1 := 10

	if i := 10; i < 11 {
		fmt.Println(i)
	}
}
