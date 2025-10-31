package main

import "fmt"

func processOp(x int, y int, op func(a int, b int)) {
	op(x, y)

}
func add(a int, b int) {
	z := a + b
	fmt.Println(z)
}
func mul(x int , y int){
	sum := x * y
	fmt.Println(sum)
}

func main() {
	processOp(3, 5, add)
	processOp(3, 5, mul)
}
