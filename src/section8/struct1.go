// 구조체 기본(1)
package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func (a Account) Calculate() float64 {
	return a.balance + (a.balance * a.interest)
}

func main() {
	// 구조체 선언 -> 구조체 인스턴스 생성

	// 예제1 - 정의된 구조체에 값을 선언하지 않으면 0으로 초기화
	kim := Account{number: "245-901", balance: 10000, interest: 0.015}
	lee := Account{number: "245-902", balance: 50000}
	park := Account{number: "245-904", interest: 0.025}
	cho := Account{"245-905", 15000, 0.025}

	fmt.Println("Ex1 : ", kim)
	fmt.Println("Ex1 : ", lee)
	fmt.Println("Ex1 : ", park)
	fmt.Println("Ex1 : ", cho)
	fmt.Println()
	// 예제2
	fmt.Println("Ex1 : ", kim.Calculate())
	fmt.Println("Ex1 : ", lee.Calculate())
	fmt.Println("Ex1 : ", park.Calculate())
	fmt.Println("Ex1 : ", cho.Calculate())
}
