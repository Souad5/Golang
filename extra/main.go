package main

import "fmt"

func call(){
	add := func(x int, y int){
		sum := x + y
		fmt.Println(sum)

	}
	add(5,6)
}

func main(){
 call()
}

func init(){
	fmt.Println("I'm run first")
}