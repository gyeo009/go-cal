package operation_test

import (
	"go-cal/internal/operation"
	"go-cal/internal/parser"
	"go-cal/internal/utils"
	"strings"
	"testing"
)

func TestConversion(t *testing.T) {
	testExpression := "1+   14-3*(4-     2/5)*7+8"
	testExpression = strings.ReplaceAll(testExpression, " ", "")
	testExpression = utils.AddSpacesBetweenOperators(testExpression)

	parsed, _ := parser.ExpressionParser(testExpression)

	answerExpression := []string{"1", "14", "+", "3", "4", "2", "5", "/", "-", "*", "7", "*", "-", "8", "+"}

	res := operation.Conversion(parsed)
	for i, v := range res {
		target := answerExpression[i]

		if strings.Compare(target, v) != 0 {
			t.Error("Conversion Function make wrong result.")
		}
	}

}

func BenchmarkConversion(b *testing.B) {
	benchmarkExpression := "1+14-3*(4-2/5)*7+8"
	benchmarkExpression = strings.ReplaceAll(benchmarkExpression, " ", "")
	benchmarkExpression = utils.AddSpacesBetweenOperators(benchmarkExpression)
	parsed, _ := parser.ExpressionParser(benchmarkExpression)

	b.ResetTimer()
	for range b.N {
		operation.Conversion(parsed)
	}

}
