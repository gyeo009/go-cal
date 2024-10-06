// 중위 표기식 입력을 후위 표기식으로 바꿔줌
// 변환 과정
/*
	1. 피연산자는 그대로 출력
	2. 연산자는 스택이 비어있으면 자신을 바로 추가
	3. stack의 top이 2의 연산자보다 우선순위가 낮을 때까지 pop, 이후 자신을 담음
	4. 단, 여는 괄호는 닫는 괄호가 나올 때까지 pop하지 않음
	5. 닫는 괄호가 나오면 여는 괄호가 나올때까지 꺼내서 출력
	6. 마지막에 도착하면 스택에서 차례로 꺼내 출력

*/
package operation

import (
	"strconv"
)

// isOperator 함수 (문자열을 기준으로 판단)
func IsOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/" || token == "(" || token == ")"
}

// isNum 함수 (문자열을 기준으로 판단)
func IsNum(token string) bool {
	_, err := strconv.ParseFloat(token, 64) // 숫자로 변환할 수 있으면 숫자
	return err == nil
}

func OperatorPriority(str string) int32 { // go에는 char는 없고 rune(int32와 동일) 혹은 byte 사용
	if str == "+" || str == "-" {
		return 1
	} else if str == "*" || str == "/" {
		return 2
	} else if str == "(" || str == ")" {
		return 0
	}

	return -1
}
