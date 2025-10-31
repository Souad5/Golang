package main

import "fmt"

func counter() (func(), func(), func()) {
	count := 0
	fmt.Println("Initial count:",count)

	increment := func() {
		count = count + 1
		fmt.Println("Incremented! Current count:",count)
	}

	decrement := func() {
		count -= 1 
		fmt.Println("Decremented! Current count:", count)
	}

	reset := func() {
		count = 0
		fmt.Println("Reset! Current count:", count)
	}

	return increment, decrement, reset
}

func called() {
	inc, dec, reset := counter()
	inc()
	inc()
	dec()
	reset()
	inc()

}