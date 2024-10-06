package operation

func Conversion(tokens []string) []string {
	var result []string // 결과를 저장할 슬라이스
	stack := []string{} // 연산자를 저장할 스택

	for _, token := range tokens {
		if IsOperator(token) {
			if token == "(" {
				stack = append(stack, token) // '('를 스택에 추가
			} else if token == ")" {
				// 스택에서 '('가 나올 때까지 pop하고 결과에 추가
				for len(stack) > 0 && stack[len(stack)-1] != "(" {
					result = append(result, stack[len(stack)-1]) // 결과에 추가
					stack = stack[:len(stack)-1]                 // 스택에서 pop
				}
				if len(stack) > 0 && stack[len(stack)-1] == "(" {
					stack = stack[:len(stack)-1] // '(' 도 pop
				}
			} else {
				// 현재 연산자의 우선순위와 스택의 top의 우선순위 비교
				cur_priority := OperatorPriority(token)
				for len(stack) > 0 && cur_priority <= OperatorPriority(stack[len(stack)-1]) {
					result = append(result, stack[len(stack)-1]) // 결과에 추가
					stack = stack[:len(stack)-1]                 // 스택에서 pop
				}
				stack = append(stack, token) // 현재 연산자를 스택에 추가
			}

		} else if IsNum(token) {
			result = append(result, token) // 숫자를 결과에 추가
		}
	}

	// 스택에 남아 있는 연산자를 모두 결과에 추가
	for len(stack) > 0 {
		result = append(result, stack[len(stack)-1]) // 결과에 추가
		stack = stack[:len(stack)-1]                 // 스택에서 pop
	}

	return result // 후위 표기식의 문자열 슬라이스를 반환
}
