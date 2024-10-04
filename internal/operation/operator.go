// 실제 연산 수행 함수
package operation

import (
	"fmt"
)

func OperateArithmetic(expression string) {
	// 1. 후위 표기식으로 변환

	// 1-1. 속도 테스트

	// 2. 연산 진행
	fmt.Println("연산중입니다...")

	fmt.Println("연산 완료, 결과값 : ")
}

func isOperator(ch byte) bool {
	if ch == '+' || ch == '-' || ch == '*' ||
		ch == '/' || ch == '(' || ch == ')' {
		return true
	} else {
		return false
	}
}

func operatorPriority(ch byte) int32 { // go에는 char는 없고 rune(int32와 동일) 혹은 byte 사용
	if ch == '+' || ch == '-' {
		return 1
	} else if ch == '*' || ch == '/' {
		return 2
	} else if ch == '(' || ch == ')' {
		return 0
	}

	return -1
}
