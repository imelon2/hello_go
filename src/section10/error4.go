// Go Error 처리 기초 (4)

package main

import (
	"errors"
	"fmt"
	"log"
)

func notZero1(n int) (string, error) {
	if n != 0 {
		s := fmt.Sprint("Hello Golang : ", n)
		return s, nil
	}

	return "", errors.New("0 를 입력했습니다. 애러 발생!")
}

func main() {
	a, err := notZero1(1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Ex1 : ", a)

	b, err := notZero1(0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Ex2 : ", b)
}
