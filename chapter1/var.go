package main

import "fmt"

var y = 43

// 함수 내부에 값을 선언하거나할 때는 단축선언 연산자를 사용하고 전역에 변수를 설정할 때는 var 를 쓰면된다.
// 바로 값을 넣지 않을 경우 타입을 명시해주면 오류가 나지않는다. // 기본값 제공한다. 정수에는 0 소수점에는 0.0 문자열은 "" 등등.
var z int

func main() {

	// 선언과 동시에 값을 넣는다
	x := 42
	fmt.Println(x)

	fmt.Println(y)
}
