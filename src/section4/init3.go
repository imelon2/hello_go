package main

import (
	"fmt"
	"section4/lib3"
)

func init() {
	fmt.Println("Init Method Start!!")
}

var num int32

func init() {
	num = 30
	fmt.Println("Main init Start")
}

func main() {
	fmt.Println("10 보다 큰 수? : ", lib3.CheckNum(num))
}
