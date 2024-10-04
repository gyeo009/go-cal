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
	"strings"
)

func Conversion(input string) string {
	var result strings.Builder
	stack := []byte{}

	for i := 0; i < len(input); i++ {
		v := input[i]
		if isOperator(byte(v)) {
			if byte(v) == '(' {
				stack = append(stack, byte(v))
				result.WriteByte(' ')
			} else if byte(v) == ')' {
				for len(stack) > 0 && stack[len(stack)-1] != '(' {
					result.WriteByte(stack[len(stack)-1])
					result.WriteByte(' ')

					stack = stack[:len(stack)-1]
				}
				stack = stack[:len(stack)-1] // '(' 도 pop
			} else {
				// 3. stack의 top이 2의 연산자보다 우선순위가 낮을 때까지 pop, 이후 자신을 담음
				cur_priority := operatorPriority(byte(v))
				for len(stack) > 0 && cur_priority <= operatorPriority(stack[len(stack)-1]) {
					result.WriteByte(stack[len(stack)-1])
					result.WriteByte(' ')
					stack = stack[:len(stack)-1]
				}

				stack = append(stack, byte(v))
			}

		} else if isNum(byte(v)) {
			for i < len(input) {
				if !isNum(input[i]) {
					break
				}
				result.WriteByte(input[i])
				i++
			}
			result.WriteByte(' ') // 피연산자 구분을 위해 공백 추가
			i--
		}
	}

	for len(stack) > 0 {
		result.WriteByte(stack[len(stack)-1])
		result.WriteByte(' ')

		stack = stack[:len(stack)-1]
	}

	return result.String()
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

func isNum(ch byte) bool {
	return 48 <= ch && ch <= 57 // 0 ... 9
}
