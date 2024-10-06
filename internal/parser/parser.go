// 수식을 파싱하는 로직
package parser

import (
	"errors"
	"strings"
)

func ExpressionParser(input string) ([]string, error) {
	if input == "" {
		return nil, errors.New("empty Expression")
	}

	tokens := tokenize(input)
	return tokens, nil
}

// 토큰화 함수
func tokenize(expr string) []string {
	// 공백을 기준을 분리, 즉 연산자와 숫자를 분리하는 간단한 토큰화
	return strings.Fields(expr)
}
