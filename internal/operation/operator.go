// 실제 연산 수행 함수
package operation

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func OperateArithmetic(input string) {
	// 1. 후위 표기식으로 변환
	postfix := MeasureExecutionTimeString(conversion, input)
	fmt.Printf("변환 완료, 결과값 : %v\n", postfix)
	// 1-1. 속도 테스트

	// 2. 연산 진행
	res := MeasureExecutionTime(Evaluate, postfix)

	fmt.Printf("연산 완료, 결과값 : %v\n", res)
}

func Evaluate(postfix string) int { // 기본 2초 이상은 걸린다고 가정
	time.Sleep(time.Second * 2)
	stack := []int{}
	tokens := strings.Fields(postfix)

	for _, token := range tokens {
		token_byte := ([]byte(token))[0]
		if isOperator(token_byte) { // 연산자인 경우 피연산자 두 개를 꺼내서 연산 진행
			num1 := stack[len(stack)-1]
			num2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			res := operation(num1, num2, token_byte)
			stack = append(stack, res)

		} else { // 피연산자인 경우 스택에 추가
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	return stack[0]
}

func operation(num1, num2 int, oper byte) int {
	if oper == '+' {
		return num2 + num1
	} else if oper == '-' {
		return num2 - num1
	} else if oper == '*' {
		return num2 * num1
	} else if oper == '/' {
		return num2 / num1
	}

	return -1
}
