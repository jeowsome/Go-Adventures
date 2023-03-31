package main

import "fmt"

type BankAccount struct {
	Balance float64
}

// Create the Deposit(), Withdraw() and DisplayBalance() methods below:
func (b *BankAccount) Deposit(amount float64) {
	b.Balance += amount
	fmt.Printf("Deposited %.2f, your new balance is: %.2f\n", amount, b.Balance)
}

func (b *BankAccount) Withdraw(amount float64) {
	if b.Balance-amount < 0 {
		fmt.Printf("Not enough funds in your bank account, funds: %.2f\n", b.Balance)
		return
	}
	b.Balance -= amount
	fmt.Printf("Withdrew %.2f, your new balance is: %.2f\n", amount, b.Balance)
}

func (b BankAccount) DisplayBalance() {
	fmt.Printf("Your current balance is: %.2f\n\n", b.Balance)
}

// DO NOT delete or modify the contents of the main() function!
func main() {
	var acc BankAccount
	for {
		displayMenu()
		var amount float64
		var choice string
		fmt.Scanln(&choice, &amount)

		switch choice {
		case "deposit":
			acc.Deposit(amount)
			fmt.Println()
		case "withdraw":
			acc.Withdraw(amount)
			fmt.Println()
		case "balance":
			acc.DisplayBalance()
		case "exit":
			fmt.Println("Exiting...")
			return
		}
	}
}

func displayMenu() {
	fmt.Println("Welcome to the bank! Enter the operation you want to perform:")
	fmt.Println("deposit | withdraw | balance | exit")
}
