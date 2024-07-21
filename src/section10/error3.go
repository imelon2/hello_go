// Go Error 처리 기초 (3)

package main

import (
	"errors"
	"fmt"
)

// error 패키지
// errors 패키지의 new 메소드를 활용한 에러 생성
// Errorf 보다 더 많이 사용
func main() {
	var err1 error = errors.New("Error Occured - 1")
	err2 := errors.New("Error occurred - 2")

	fmt.Println("error1 : ", err1 /* == err1.Error()*/)
	fmt.Println("error2 : ", err2 /* == err2.Error()*/)

}
