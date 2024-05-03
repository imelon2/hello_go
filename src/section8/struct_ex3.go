package main

import (
	"fmt"
)

type Account1 struct {
	number   string
	balance  float64
	interest float64
}

func (a *Account1) CalcuateD(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

func (a Account1) CalcuateF(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

func main() {
	// 예제1
	kim := Account1{"245-901", 1000, 0.015}
	lee := Account1{"245-902", 5000, 0.025}

	fmt.Println("Ex1 : ", kim)
	fmt.Println("Ex1 : ", lee)
	fmt.Println()

	kim.CalcuateD(10000)
	lee.CalcuateF(10000)

	fmt.Println("Ex2 : ", kim)
	fmt.Println("Ex2 : ", lee)
	fmt.Println()
}
