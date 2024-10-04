// 연산식 하나에 대한 속도 평가
package operation

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

func MeasureExecutionTime(f func(string) int, postfix string) int {
	start := time.Now()                        // 함수 실행 시작 시간
	start_goroutines := runtime.NumGoroutine() // 시작 시점의 Goroutine 수

	// 함수 실행
	res := f(postfix)
	fmt.Println("\n\n")
	elapsed := time.Since(start).Milliseconds()

	// r := 3475393 / 234 * 1242312 * 1241235 / 12356123
	// fmt.Print(r, "\n")

	// 종료 후 경과 시간 계산
	elapsed2 := time.Since(start).Milliseconds()
	end_goroutines := runtime.NumGoroutine() // 종료 시점의 Goroutine 수
	numCPU := runtime.NumCPU()               // CPU 수
	numThreads := runtime.GOMAXPROCS(0)      // 사용 중인 최대 스레드 수

	// 함수 이름 출력
	funcValue := reflect.ValueOf(f)
	if funcValue.Kind() == reflect.Func {
		pc := runtime.FuncForPC(funcValue.Pointer())
		fmt.Printf("Function name: %s\n", pc.Name()) // 함수 이름 출력
	}
	// 결과 출력
	fmt.Printf("Execution time: %d ms\n", elapsed)
	fmt.Printf("Execution time2: %d ms\n", elapsed2)

	fmt.Printf("Start Goroutines: %d, End Goroutines: %d\n", start_goroutines, end_goroutines)
	fmt.Printf("Available CPUs: %d, Max Threads (GOMAXPROCS): %d\n", numCPU, numThreads)

	return res
}
func MeasureExecutionTimeString(f func(string) string, input string) string {
	start := time.Now()                        // 함수 실행 시작 시간
	start_goroutines := runtime.NumGoroutine() // 시작 시점의 Goroutine 수

	// 함수 실행
	res := f(input)
	elapsed := time.Since(start).Milliseconds()

	// 종료 후 경과 시간 계산
	end_goroutines := runtime.NumGoroutine() // 종료 시점의 Goroutine 수
	numCPU := runtime.NumCPU()               // CPU 수
	numThreads := runtime.GOMAXPROCS(0)      // 사용 중인 최대 스레드 수

	// 함수 이름 출력
	funcValue := reflect.ValueOf(f)
	if funcValue.Kind() == reflect.Func {
		pc := runtime.FuncForPC(funcValue.Pointer())
		fmt.Printf("Function name: %s\n", pc.Name()) // 함수 이름 출력
	}
	// 결과 출력
	fmt.Printf("Execution time: %d ms\n", elapsed)

	fmt.Printf("Start Goroutines: %d, End Goroutines: %d\n", start_goroutines, end_goroutines)
	fmt.Printf("Available CPUs: %d, Max Threads (GOMAXPROCS): %d\n", numCPU, numThreads)

	return res
}
