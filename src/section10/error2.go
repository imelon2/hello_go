// Go Error 처리 기초 (2)

package main

import (
	"fmt"
	"log"
)

func notZero(n int) (string, error) {
	if n != 0 {
		s := fmt.Sprint("Hello Golang : ", n)
		return s, nil
	}

	return "", fmt.Errorf("%d를 입력했습니다. 애러 발생!", n)
}

func main() {
	a, err := notZero(1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Ex1 : ", a)

	b, err := notZero(0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Ex2 : ", b)
}
