package main

import (
	"fmt"
	checkUp "section4/lib"
	"section4/lib2"
)

func main() {
	fmt.Println("100보다 큰 수? : ", checkUp.CheckNum(101))
	fmt.Println("100보다 큰 수? : ", lib2.CheckNum1(101))
	fmt.Println("1000보다 큰 수? : ", lib2.CheckNum2(999))
}