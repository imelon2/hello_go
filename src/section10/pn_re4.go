// Panic & Recover (4) 최종
package main

import (
	"fmt"
	"os"
)

func fileOpen(filename string) {
	// defer 함수 (panic 호출되면 실행)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("File Open Error : ", r)
		}
	}()
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Ex1 : ", f.Name())
	}

	defer f.Close() // Resource 돌려주는 작업
}

func main() {
	fileOpen("undefined.txt")
	fmt.Println("End Main")
}
