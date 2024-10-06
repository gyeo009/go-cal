package operation_test

import (
	operation "go-cal/internal/operation"
	"strings"
	"testing"
)

func TestConversion(t *testing.T) {
	testExpression := "1+14-3*(4-2/5)*7+8"
	answerExpression := "1 14 + 3  4 2 5 / - * 7 * - 8 + "
	res := operation.Conversion(testExpression)

	if strings.Compare(res, answerExpression) != 0 {
		t.Error("Conversion Function make wrong result.")
	}
}

func BenchmarkConversion(b *testing.B) {
	benchmarkExpression := "1+14-3*(4-2/5)*7+8"
	b.ResetTimer()
	for range b.N {
		operation.Conversion(benchmarkExpression)
	}

}
