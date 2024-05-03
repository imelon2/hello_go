package main

import (
	"fmt"
)

type Account struct {
	number   string
	balance  float64
	interest float64
}

// 생성자를 위한 함수 구현 - Go의 관례
func NewAccount(number string, balance, interest float64) *Account {
	return &Account{number, balance, interest}
}

func main() {
	// 구조체 생성자 패턴 예제

	// 예제1
	kim := Account{"245-901", 1000000, 0.015} // 생성자 초기화

	var lee *Account = new(Account) // 이건 생성자 안됨

	fmt.Println("Ex1 : ", kim)
	fmt.Println("Ex1 : ", lee)

	// 예제2
	park := NewAccount("245-903", 170000, 0.04)
	fmt.Println("Ex2 : ", park)

}
