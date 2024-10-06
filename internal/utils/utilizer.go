package utils

import (
	"fmt"
	"go-cal/internal/operation"
	"strings"
)

func AddSpacesBetweenOperators(expression string) string {
	var result strings.Builder

	for i := 0; i < len(expression); i++ {
		result.WriteByte(expression[i]) // 현재 문자 추가

		// 다음 문자가 연산자일 때 지금 공백 추가
		if i < len(expression)-1 && operation.IsOperator(string(expression[i+1])) {
			result.WriteByte(' ')
		}

		// 현재 연산자면서 다음 문자가 숫자면 지금 공백 추가
		if i < len(expression)-1 && operation.IsOperator(string(expression[i])) && !operation.IsOperator(string(expression[i+1])) {
			result.WriteByte(' ')
		}
	}

	fmt.Println(result.String())

	return result.String()
}
