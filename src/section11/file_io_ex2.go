// 파일 I/O (2)

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}

// 파일 읽기, 버퍼 사용 쓰기 -> bufio 패키지 활용
// ioutil, bufio 등은 io.Reader, io.Writer 인터페이스를 구현함
// 즉, Writer, Read 메소드 사용 가능
/*
	type Reader interface {
		Read(p []byte) (n int, err error)
	}
	type Writer interface {
		Writer(p []byte) (n int, err error)
	}
*/

/*
	권한은 상수로 정의되어 있음
	https://pkg.go.dev/os@go1.22.5#pkg-constants
	const (
	// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
	O_RDONLY int = syscall.O_RDONLY // open the file read-only.
	O_WRONLY int = syscall.O_WRONLY // open the file write-only.
	O_RDWR   int = syscall.O_RDWR   // open the file read-write.
	// The remaining values may be or'ed in to control behavior.
	O_APPEND int = syscall.O_APPEND // append data to the file when writing.
	O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
	O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
	O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
	O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.

)
*/
func main() {
	// file이 없으면 만들고, file이 있으면 읽고 쓰기
	file, err := os.OpenFile("test_write2.txt", os.O_CREATE|os.O_RDWR, os.FileMode(0777))
	errCheck(err)

	// bufio 파일 쓰기 예제
	wt := bufio.NewWriter(file)
	wt.WriteString("(1) 안녕안녕 나는 지수야 헬륨가스 먹었더니 요러케됬지 ~\n")
	wt.Write([]byte("(2) 안녕안녕 나는 지수야 헬륨가스 먹었더니 요러케됬지 ~\n"))

	// 이전까지는 Buffer에 데이터가 쌓임
	fmt.Printf("사용중인 Buffer Size (%d byte)\n", wt.Buffered())
	fmt.Printf("남은 Buffer Size (%d byte)\n", wt.Available())
	fmt.Printf("전체 Buffer Size (%d byte)\n", wt.Size())

	// 버퍼의 내용을 디스크에 기록 하는 함수 실행 -> 버퍼는 비워짐
	wt.Flush()

	fmt.Println("==========================================")
	rt := bufio.NewReader(file)
	fi, _ := file.Stat()
	b := make([]byte, fi.Size())

	file.Seek(0, io.SeekCurrent)
	data, _ := rt.Read(b) // 읽기

	fmt.Println(data)
}
