// // 사용자 패키지 작성 및 문서화 예제
// package main

// import (
// 	"fmt"
// 	oper "section12/arithmetic" // alias
// )

// `go mod init section12`로 .mod 생성 필수!

// 사용자 패키지 작성 & Go문서화
// 기준은 GOPATH/src
// 패키지 폴더명(디렉토리 명) 명확하게 지정
// 패키지 파일의 package `이름(ex. package arithmetic)`으로 사용한다.
// package main을 제외하고 package 문서에 등록
// 기본적으로 GOROOT의 패키지(pkg)를 검색 후, GOPATH의 패키지(pkg) 검색
// go install 명령어 실행의 경우 -> GOPATH/pkg에 등록
// godic -http=:PORT -> pkg이동 후, 본인 패키지 메소드 및 주석 확인
// func main() {
// 	nums := oper.Numbers{100, 10}
// 	fmt.Println("Package Used(1): ", nums.Plus())
// 	fmt.Println("Package Used(2): ", nums.Minus())
// 	fmt.Println("Package Used(3): ", nums.Multi())
// 	fmt.Println("Package Used(4): ", nums.Divide())
// 	fmt.Println("Package Used(5): ", nums.SquareMinus())
// 	fmt.Println("Package Used(6): ", nums.SquarePlus())
// 	fmt.Println("Package Used(7): ", nums.X)
// 	fmt.Println("Package Used(8): ", nums.Y)
// }
