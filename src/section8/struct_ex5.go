package main

import "fmt"

// 직원
type Employee struct {
	name   string
	salary float64
	bonus  float64
}

// 임원
type Executive struct {
	employee     Employee // 구조체 임베디드
	specialBonus float64
}

// 원래 메소드
func (e Employee) Calcuate() float64 {
	return e.salary + e.bonus
}

// Overriding
func (e Executive) Calcuate() float64 {
	return e.employee.salary + e.employee.bonus + e.specialBonus
}

// 구조체 오버라이딩 패턴
func main() {
	// 예제1
	// 직원
	ep1 := Employee{"kim", 2000, 150}
	ep2 := Employee{"park", 5000, 200}

	// 임원
	ex1 := Executive{
		Employee{"lee", 9000, 250},
		1000,
	}

	fmt.Println("Ex1 : ", ep1.Calcuate())
	fmt.Println("Ex1 : ", ep2.Calcuate())

	// Executive 구조체 호출
	fmt.Println("Ex2 : ", ex1.employee.Calcuate())
	fmt.Println("Ex2 : ", ex1.Calcuate())
}
