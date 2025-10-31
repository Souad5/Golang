package main

import (
	"fmt"
	"new/mathlib"
)



func num(num1 int, num2 int) (int,int,float32) {
	sum := num1 + num2
	mul := num1 * num2
	div :=float32 (num1)/ float32 (num2)
	return sum, mul,div
}
func strings(name string){
	fmt.Println(`Learning Go lang with Habib`, name)
}


func mul (num1 int, num2 int)int{
	sum := num1 * num2
	return sum 
}
func main() {
	// a := 15
	// b := 10
	// sum, mul,div := num(a, b)
	// fmt.Println(sum,mul,div)
	// strings("Souad")
	wel()
	var name =name()
	fmt.Println(name)
	num1,num2 :=sum()
	new:= mul(num1,num2)
	fmt.Println(new)
	mathlib.Sum()
	newww()
}
