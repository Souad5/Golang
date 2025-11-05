package main

import "fmt"
func numbers(num *[3]int){
fmt.Println(num)
}

func main(){
	a := 10
		fmt.Println(a)

	p:= &a
	*p = 20
	fmt.Println(a)
	num := [3] int{1,2,3}
	numbers(&num)
}