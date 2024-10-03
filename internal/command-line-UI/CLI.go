package commandLineUI

import (
	"fmt"
	"time"
)

var option_string string = "===============================================================================================================\n" +
	`
	1. 사칙연산(정수)
	2. 사칙연산(소수점이 있는 정수)
	3. 이전 기록 불러오기
	4. 만든 사람
	5. 종료...
` + "\n===============================================================================================================\n"

var about string = `
This project was created by a senior computer science student who is currently exploring the Go programming language.
Having started learning Go just three days ago, this project represents an early yet enthusiastic dive into the world of concurrent programming,
with a focus on building efficient and scalable applications.

Despite being new to Go, I have a strong foundation in software development from my undergraduate studies and am passionate about continuously learning new technologies.
This project reflects both my learning journey and my desire to grow as a developer.

이 프로젝트는 Go 프로그래밍 언어를 탐색하고 있는 컴퓨터공학과 4학년 학생에 의해 만들어졌습니다.
Go를 배우기 시작한 지 3일밖에 되지 않았지만, 이 프로젝트는 효율적이고 확장 가능한 애플리케이션을 구축하는 데 중점을 둔 초기 단계의 열정적인 시도를 나타냅니다.

Go에 대해서는 아직 많이 배우고 있지만, 학부 과정을 통해 소프트웨어 개발에 대한 튼튼한 기초를 갖추고 있으며,
새로운 기술을 지속적으로 배우는 것에 대한 열정이 있습니다.
이 프로젝트는 제 학습 여정을 반영하며, 개발자로서 성장하고자 하는 저의 의지를 보여줍니다.
`

func CommandLineInterface() {
	// ASCII Art 출력
	fmt.Println(PrintAsciiArt() + "\n\n")

	// 명령 옵션 출력
	for {
		fmt.Println("\n\n" + option_string)
		fmt.Print("실행할 명령 번호를 입력하세요... : ")

		// 인자 받기
		var input string
		fmt.Scanln(&input)

		switch input {
		case "1": // 사칙연산(정수)
			func() {

			}()
		case "2": // 사칙연산(소수점이 있는 정수)
			func() {

			}()
		case "3": // 사칙연산(이전 기록 불러오기), UTC 표준 시간도 기록하자
			func() {

			}()
		case "4": // 만든 사람
			fmt.Println(about)
		case "5":
			fmt.Println("이용해주셔서 감사합니다.")
			return
		default:
			fmt.Println("1부터 5 사이의 번호를 입력하세요...")
			time.Sleep(2 * time.Second)
		}
	}

}
