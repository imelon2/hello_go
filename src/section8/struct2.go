// 구조체 기본(2)
package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func (a *Account) Calculate() float64 {
	return a.balance + (a.balance * a.interest)
}

func main() {
	// 다양한 선언 방법
	// &struct 포인터를 받아고기,,, 역참조를 또 하기에 속도가 느리다
	// 하지만 인터페이스 메소드를 선언만 해둔 후, Overriding해서 메서드에 포잉ㄴ터 리시버를 사용하는 경우 포인터를 사용해야한다.

	// 예제1
	// 선언 방법 1
	var kim = new(Account) // kim type = *Account
	kim.number = "245-901"
	kim.balance = 10000
	kim.interest = 0.015

	// 선언 방법 2
	hong := &Account{"245-902", 15000000, 0.04}

	// 선언 방법 3
	lee := new(Account)
	lee.number = "245-903"
	lee.balance = 2000
	lee.interest = 0.5

	fmt.Println("Ex1 : ", kim)
	fmt.Println("Ex1 : ", hong)
	fmt.Println("Ex1 : ", lee)

	fmt.Printf("Ex2 : %v \n", kim)
	fmt.Printf("Ex2 : %v \n", hong)
	fmt.Printf("Ex2 : %v \n", lee)

	// 예제2
	fmt.Println("Ex3 : ", kim.Calculate())
	fmt.Println("Ex3 : ", hong.Calculate())
	fmt.Println("Ex3 : ", lee.Calculate())
}
