package main

import "fmt"

var a = 10
var b = 30
func main() {
	fmt.Println(b)
}

func init(){
	fmt.Println(a)
	b = 45
}