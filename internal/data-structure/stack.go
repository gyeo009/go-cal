// go에서 stack은 slice로 충분히 구현할 수 있기 때문에 stack 표준 자료형이 존재하지않는다.
// 그래서 이 type도 사용하지는 않을 거지만, 공부를 위해 구현
package datastructure

import "fmt"

// []interface{} 는 interface{} type의 요소를 가지는 슬라이스를 의미
// 즉 다양한 type의 값을 저장할 수 있는 동적 배열을 제공
// 따라서 Stack은 interface가 아니라 Slice이다
// 내부에 메소드를 정의할 수 없다
type Stack []interface{}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(val interface{}) {
	*s = append(*s, val) // go는 슬라이스에 요소를 추가하면 새로운 슬라이스를 만들어서, 기존 s에 재할당해야 함
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		fmt.Println("Stack is empty!")
		// 에러를 발생시켜야 한다
	} else {
		top := len(*s) - 1
		*s = (*s)[:top] // 스택 마지막 요소 제거 후 s에 재할당
	}
}

func (s Stack) TopElem() interface{} {
	if s.IsEmpty() {
		fmt.Println("Stack is empty!")
		// 에러를 발생시켜야 한다
	}

	return s[len(s)-1]
}
