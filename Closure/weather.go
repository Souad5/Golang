package main

import "fmt"

func temperature(celsius int) (func(), func(int)) {
	fmt.Println("Initial Temperature:", celsius, "°C")

	toFahrenheit := func() {
		// TODO: convert Celsius to Fahrenheit and print
		fah := (float64(celsius) * 1.8)  + 32
		fmt.Println("Initial Temperature:",fah,"°C")
	}

	increaseTemp := func(delta int) {
		// TODO: add delta to Celsius and print new temperature
		celsius += delta
		fmt.Println(celsius) 
	}

	return toFahrenheit, increaseTemp
}

func alu() {
	toF, inc := temperature(25)
	toF()
	inc(5)
	toF()
}
