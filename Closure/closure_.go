package main

import "fmt"

func bankAccount(initialBalance int) (func(int), func(int)) {
	balance := initialBalance
	fmt.Println("Initial balance:", balance)

	deposit := func(amount int) {
		balance += amount
		fmt.Println("Deposited", amount)

		fmt.Println("New balance:", balance)
	}
	withdraw := func(amount int) {
		if amount > balance {
			fmt.Println("Insufficient funds!")
		} else {
			balance -= amount
			fmt.Println("Withdrew:", amount)
			fmt.Println("New balance:", balance)
		}
	}
	return deposit, withdraw
}

func amount() {
	d, withdraw := bankAccount(1000)
	d(500)
	withdraw(200)
	withdraw(2000)
	d(100)
}
