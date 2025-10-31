package main
import "fmt"

func wel(){
	fmt.Println("Welcome to new app")
}
func name()string {
	var name string
	fmt.Println("Enter your name - ")

	fmt.Scanln(&name)
	return name
}

func sum ()(int,int){
	var num1 int
	var num2 int
	fmt.Println("Enter a number 1 ---")
	fmt.Scanln(&num1)
	fmt.Println("Enter a number 2 ---")
	fmt.Scanln(&num2)
	return num1 , num2
}

