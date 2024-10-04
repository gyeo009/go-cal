// go에서 stack은 slice로 충분히 구현할 수 있기 때문에 stack 표준 자료형이 존재하지않는다.
// 그래서 이 type도 사용하지는 않을 거지만, 공부를 위해 구현
package datastructure

type Stack []interface{}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
