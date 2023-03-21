package main

import "fmt"

var y = 42

// 타입을 아예 지정해 줄 경우 정적 프로그래밍이다.
// 동적 프로그래밍이 아니다.

var z string = "손을 흔들어 봅시다."
var a string = `나는 손을 흔들어 봅시다.`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
