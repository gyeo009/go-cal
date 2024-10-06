// 수식을 계산하는 로직
package evaluator

import (
	"go-cal/internal/parallel"
	"go-cal/internal/parser"
)

func EvaluateExpression(expression string) (float64, error) {
	tokens, err := parser.ExpressionParser(expression)
	if err != nil {
		return 0, err
	}

	// 병렬 처리를 위한 분기
	if shouldParallelize(tokens) {
		return parallel.EvaluateInParallel(tokens)
	} else {
		return parallel.SimpleEvaluate(tokens), nil
	}
}

func shouldParallelize(tokens []string) bool {
	multiplierDividerCount := 0
	for _, token := range tokens {
		if token == "*" || token == "/" {
			multiplierDividerCount++
		}
	}
	return multiplierDividerCount >= 20
}
