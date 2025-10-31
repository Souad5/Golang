package mathlib

import "fmt"

var a = 3
var b = 5

func math (x int, y int) int{
	sum := x + y
	return sum
}

func Sum() {
	var sum = math(a,b )
	fmt.Println(sum)
}