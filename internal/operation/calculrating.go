// 실제 연산 수행 함수
package operation

import (
	"strconv"
	"time"
)

// func OperateArithmetic(input []string) {
// 	// 1. 후위 표기식으로 변환
// 	postfix := Conversion(input)
// 	fmt.Printf("변환 완료, 결과값 : %v\n", postfix)

// 	// 2. 연산 진행
// 	res := Evaluate(postfix)
// 	fmt.Printf("연산 완료, 결과값 : %v\n", res)
// }

func Evaluate(postfix []string) int { // 입력을 []string으로 변경
	time.Sleep(time.Second * 2) // 기본 2초 이상 걸린다고 가정
	stack := []int{}            // 연산을 위한 스택 초기화

	for _, token := range postfix {
		if IsOperator(token) { // 연산자인 경우
			num1 := stack[len(stack)-1]         // 스택에서 마지막 피연산자 꺼내기
			num2 := stack[len(stack)-2]         // 스택에서 두 번째 피연산자 꺼내기
			stack = stack[:len(stack)-2]        // 스택에서 피연산자 두 개 제거
			res := operation(num1, num2, token) // 연산 수행
			stack = append(stack, res)          // 결과를 스택에 추가

		} else { // 피연산자인 경우
			num, _ := strconv.Atoi(token) // 문자열을 정수로 변환
			stack = append(stack, num)    // 스택에 피연산자 추가
		}
	}

	return stack[0] // 최종 결과 반환
}

func operation(num1, num2 int, oper string) int {
	if oper == "+" {
		return num2 + num1
	} else if oper == "-" {
		return num2 - num1
	} else if oper == "*" {
		return num2 * num1
	} else if oper == "/" {
		return num2 / num1
	}

	return -1
}
