// 두 숫자의 사칙연산 계산 제공 패키지(1)
package arithmetic

// X, Y 2개의 Int 구조체
type Numbers struct {
	X int
	Y int
}

// x, y 합(+)을 계산해서 반환
func (o *Numbers) Plus() int {
	return o.X + o.Y
}

// x, y 차(-)을 계산해서 반환
func (o *Numbers) Minus() int {
	return o.X - o.Y
}

// x, y 곱(*)을 계산해서 반환
func (o *Numbers) Multi() int {
	return o.X * o.Y
}

// x, y 나누기(/)을 계산해서 반환
func (o *Numbers) Divide() int {
	return o.X / o.Y
}
