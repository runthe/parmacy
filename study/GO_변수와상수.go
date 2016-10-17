package main

import "fmt"

var (
	str string
	num int
)

const (
	Apple = iota
	Grape
	Orage
)

func main() {

	/**
	 변수를 선언하지 않으면 변수는 초기값이 숫자형은 = 0
	 bool 타입에는 false string 형에는 "" 빈문자열을 할당한다
	 */

	var i, j, k int = 1, 2, 3
	fmt.Println(num, str)
	fmt.Println(i, j, k)
	fmt.Println(Apple, Grape, Orage)

	/**
	 암묵적 타입 변환은 안일어나기 때문에 묵시적으로 명시해주어야 한다
	 타입(값)
	 */

	var p int = 100
	var u uint = uint(p)
	var f float32 = float32(u)
	fmt.Println(f, u)

	str0 := "ABC"
	bytes := []byte(str0)
	str1 := string(bytes)
	fmt.Println(bytes, str1)

}