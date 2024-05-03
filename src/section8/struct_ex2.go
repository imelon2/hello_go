package main

import (
	"fmt"
)

type Account struct {
	number   string
	balance  float64
	interest float64
}

func CalcuateD(a Account) {
	a.balance = a.balance + (a.balance * a.interest)
}

func CalcuateP(a *Account) {
	a.balance = a.balance + (a.balance * a.interest)
}

func main() {
	// 예제1
	kim := Account{"245-901", 1000, 0.015}
	lee := Account{"245-902", 5000, 0.025}

	fmt.Println("Ex1 : ", kim)
	fmt.Println("Ex1 : ", lee)
	fmt.Println()

	CalcuateD(kim)
	CalcuateP(&lee)

	fmt.Println("Ex2 : ", kim)
	fmt.Println("Ex2 : ", lee)
	fmt.Println()
}
