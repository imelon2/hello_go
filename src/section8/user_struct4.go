// 사용자 정의 타입(4)
package main

import "fmt"

type shoppingBasket struct{ cnt, price int }

// 결제함수
func (b shoppingBasket) purchase() int {
	return b.cnt * b.price
}

// 참조 형식(얕은복사) - 원본이 수정됨
func (b *shoppingBasket) purchaseP(cnt, price int) {
	b.cnt += cnt
	b.price += price
}

// 값 전달 형식(깊은복사) - 원본이 수정x
func (b shoppingBasket) purchaseD(cnt, price int) {
	b.cnt += cnt
	b.price += price
}

func main() {
	// 구조체와 메소드 바인딩은 구조체에 포인터 타입(*)이 정의되어 있으면
	// 자동으로 참조 복사가 진행됨

	// 예제1 - 함수 호출
	bs1 := shoppingBasket{3, 5000}
	fmt.Println("Ex1 (total price) : ", bs1.purchase())

	// 예제2 - 포인터 함수 호출
	bs1.purchaseP(7, 5000)
	fmt.Println("Ex2 (total price) : ", bs1.purchase())

	// 예제3
	bs1.purchaseD(10, 5000)
	fmt.Println("Ex3 (total price) : ", bs1.purchase())
}
