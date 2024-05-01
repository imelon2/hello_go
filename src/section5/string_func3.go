package main

import (
	"fmt"
	"strings"
)

func main() {
	// 예제1 결합 - 일반연산
	str1 := "디지털 서명의 과정을 추가로 설명드리자면, A라는 사람이 '개인키'를 이용해서 내가 만든 데이터에 서명을 하고, 데이터를 검증하는 사람은 A라는 사람의 '공개키'를 통해 진짜로 A라는 사람이 서명한게 맞는지 검증을 합니다." +
		"개인키와 공개키를 통해 위와 같은 검증과정을 만들 수 있는 이유는 바로 '타원 곡선 함수' 때문입니다." +
		"타원 곡선 함수는 생성점 G(Generator Point:생성점)에 따라 활용도가 달라집니다."

	str2 := "우리가 서명을 위해 사용하는 개인키의 생성 원리를 잠깐 알아보겠습니다."

	fmt.Println("ex1 : ", str1+str2)

	// 예제2 결합 - join
	// 성능적으로 일반연산(+)보단 좋다고함
	strSet := []string{} // 슬라이스 선언
	strSet = append(strSet, str1)
	strSet = append(strSet, str2)

	fmt.Println("ex2 : ", strings.Join(strSet, ""))
}
