package main

import "fmt"

const a = 100

var b = 20

func outer() func() {
	money := 10
	fmt.Println(money)

	sum := func() {
		money = money + a + b
		fmt.Println(money)
	}
	return sum

}

func main() {
	// call := outer()
	// call()
	// call()
	// call()
	fmt.Println("Hello")
	// amount()
	// called()
	alu()
}
