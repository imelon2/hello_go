package lib3

import "fmt"

func init() {
	fmt.Println("init lib3")
}

func CheckNum(c int32) bool {
	return c > 100
}
