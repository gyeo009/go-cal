// // 병렬 처리 후
// // channel을 통해 연산 결과값 return

// // 기존: O(n)
// // 최선: O(n/p) 기대
// // 평균적: 병렬 처리 오버헤드 고려해 O(n)
// // 최악: 더 느려질 수 있다.

// // 병렬 처리할 분기 기준:
// // 1. 괄호 내부 연산식을 분기하면 안된다
// // 2. 곱셈 or 나눗셈 연산자가 20개 이상 존재할 때, 더하기나 빼기 연산자가 나오면 분기하기

// Algorithm Parallel_Evaluate:
//     // 입력 및 출력 정의
//     Input: Arithmetic expression  // 입력: 사칙연산 표현식 (문자열)
//     Output: Evaluation result  // 출력: 최종 계산 결과 (문자열 형태)

//     // 필요한 변수를 정의
//     IsInBracket: int  // 괄호 안에 있는지 확인하는 변수
//     multiplier_or_divider_count: int  // 곱셈/나눗셈 연산자가 몇 번 나왔는지 카운트하는 변수
//     current_expression_batch: string  // 현재 처리 중인 연산식 조각

//     // 초기화
//     IsInBracket = 0  // 괄호 상태 초기화 (0이면 괄호 외부)
//     multiplier_or_divider_count = 0  // 곱셈/나눗셈 연산자 카운트 초기화
//     current_expression_batch = ''  // 현재 연산식 조각 초기화

//     // 병렬 처리된 결과들을 저장할 배열
//     to_be_merged_digit_array: []string = nil  // 연산 결과가 저장될 배열
//     to_be_merged_operator_array: []string = nil  // 연산자들이 저장될 배열

//     // 최종으로 축약된 연산식을 저장할 문자열
//     reduced_expression: string = ''  // 축약된 최종 연산식

//     // 고루틴에서 사용할 변수들 초기화
//     parallelized_number: int = 0  // 병렬 처리된 연산의 고유 ID를 위한 카운터
//     result_channel: chan string = make(chan string)  // 고루틴에서 연산 결과를 받을 채널

//     // 메인 루프: 입력된 연산식을 한 글자씩 순차적으로 처리
//     i: int = 0  // 입력 인덱스 초기화
//     while i < len(Input) do {
//         current_elem = Input[i]  // 현재 처리 중인 문자

//         // 현재 문자가 연산자인지 확인
//         if current_elem.IsOperator() then
//             // 곱셈(*) 또는 나눗셈(/)일 경우 카운트 증가
//             if current_elem == '*' or current_elem == '/' then
//                 multiplier_or_divider_count++
//             // 열린 괄호일 경우 괄호 카운트 증가
//             else if current_elem == '(' then
//                 IsInBracket++
//             // 닫힌 괄호일 경우 괄호 카운트 감소
//             else if current_elem == ')' then
//                 IsInBracket--

//         // 병렬 처리 분기 조건 확인
//         if current_elem.IsOperator() && ( current_elem == '+' || current_elem == '-' ) && IsInBracket == 0 && multiplier_or_divider_count >= 20 then
//             // 병렬 처리를 위한 고루틴 생성
//             go func(expr string) {
//                 res := Evaluate(current_expression_batch)  // 현재 연산식 조각을 평가
//                 result_channel <- res  // 채널에 결과를 보냄
//             }(current_expression_batch)  // 고루틴에 현재 연산식 조각 전달

//             // 연산자를 저장하고 다음 연산을 위해 초기화
//             to_be_merged_operator_array.append(current_elem)  // 연산자를 저장
//             current_expression_batch = ''  // 현재 연산식 조각 초기화
//             multiplier_or_divider_count = 0  // 카운터 초기화
//             i++  // 인덱스 증가 후 다음 문자로 이동
//             continue  // 루프의 나머지 부분 건너뛰기

//         // 현재 문자가 연산자가 아니거나 분기 조건을 만족하지 않을 때
//         current_expression_batch += current_elem  // 현재 문자를 연산식에 추가
//         i++  // 인덱스 증가

//     // 고루틴의 결과 수집
//     for j := 0; j < parallelized_number; j++ {
//         res := <-result_channel  // 채널에서 고루틴의 결과를 받음
//         to_be_merged_digit_array = append(to_be_merged_digit_array, res)  // 결과를 저장
//     }

//     // 최종적으로 축약된 연산식을 조립
//     for i, v := range to_be_merged_digit_array {
//         reduced_expression += v  // 계산된 값을 추가
//         if !to_be_merged_operator_array.isEmpty() {  // 연산자가 남아 있을 경우
//             reduced_expression += to_be_merged_operator_array.front()  // 연산자를 추가
//             to_be_merged_operator_array.pop_front()  // 연산자를 제거
//         }
//     }

//     // 최종 연산식을 반환
//     return reduced_expression  // 축약된 연산 결과 반환

// 병렬 처리 로직
package parallel

import (
	"fmt"
	"go-cal/internal/operation"
	"sync"
)

func EvaluateInParallel(tokens []string) (float64, error) {
	var wg sync.WaitGroup
	results := make(chan float64, len(tokens))

	// 토큰을 처리할 고루틴 분배
	for i := 0; i < len(tokens); i += 20 {
		wg.Add(1)
		go func(subTokens []string) {
			defer wg.Done()
			result := SimpleEvaluate(subTokens)
			results <- result
		}(tokens[i:min(i+20, len(tokens))])
	}

	// 모든 고루틴이 끝날 때까지 기다림
	wg.Wait()
	close(results)

	// 결과 합산
	total := 0.0
	for res := range results {
		total += res
	}

	return total, nil
}

func SimpleEvaluate(tokens []string) float64 {
	// 1. 후위 표기식으로 변환
	postfix := operation.Conversion(tokens)
	fmt.Printf("변환 완료, 결과값 : %v\n", postfix)

	// 2. 연산 진행
	res := operation.Evaluate(postfix)
	fmt.Printf("연산 완료, 결과값 : %v\n", res)
	return 0.0
}
