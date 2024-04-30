package main

import "fmt"

func main() {
	// example 1
	sum1 := 0
	for i := 0; i <= 100; i++ {
		sum1 += 1
	}

	fmt.Println("ex1 : ", sum1)

	// example 2
	sum2, i := 0, 0
	for i <= 100 {
		sum2 += i // 0+1+2+3+...
		i++       // go에서 후치연산은 반환(return)이 없습니다
	}
	fmt.Println("ex2 : ", sum2)

	// example 3
	sum3, i := 0, 0
	for {
		if i > 100 {
			break
		}

		sum3 += i
		i++
	}

	fmt.Println("ex3 : ", sum3)

	// example 4
	for i, j := 0, 0; i <= 10; i, j = i+1, j+10 {
		fmt.Println("ex4 : ", i, j)
	}
}
