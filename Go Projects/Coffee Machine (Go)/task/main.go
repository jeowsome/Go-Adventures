package main

import (
	"fmt"
)

func getIngredients(coffeeType string) [4]int {
	var ingredients [4]int

	switch coffeeType {
	case "1":
		ingredients = [4]int{250, 0, 16, 4}
	case "2":
		ingredients = [4]int{350, 75, 20, 7}
	case "3":
		ingredients = [4]int{200, 100, 12, 6}
	}
	return ingredients
}

func buy(water, milk, beans, cups, money **int) {
	var coffeeType string
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu:")
	fmt.Scan(&coffeeType)

	if coffeeType == "back" {
		return
	}
	ingredients := getIngredients(coffeeType)
	makeCoffee(ingredients, *water, *milk, *beans, *cups, *money)
}

func makeCoffee(ingredients [4]int, water, milk, beans, cups, money *int) {
	switch {
	case *water-ingredients[0] < 0:
		fmt.Println("Sorry, not enough water!")
	case *milk-ingredients[1] < 0:
		fmt.Println("Sorry, not enough milk!")
	case *beans-ingredients[2] < 0:
		fmt.Println("Sorry, not enough beans!")
	case *cups < 0:
		fmt.Println("Sorry, not enough cups!")
	default:
		fmt.Println("I have enough resources, making you a coffee!")
		*water -= ingredients[0]
		*milk -= ingredients[1]
		*beans -= ingredients[2]
		*money += ingredients[3]
		*cups--
	}
}

func fill(water, milk, beans, cups **int) {
	var addedWater, addedMilk, addedBeans, addedCups int
	fmt.Println("Write how many ml of water you want to add:")
	fmt.Scan(&addedWater)
	fmt.Println("Write how many ml of milk you want to add:")
	fmt.Scan(&addedMilk)
	fmt.Println("Write how many grams of coffee beans you want to add:")
	fmt.Scan(&addedBeans)
	fmt.Println("Write how many disposable cups you want to add: ")
	fmt.Scan(&addedCups)

	**water += addedWater
	**milk += addedMilk
	**beans += addedBeans
	**cups += addedCups
}

func take(money **int) {
	fmt.Printf("I gave you $%d\n", *money)
	fmt.Println("")
	**money = 0
}

func displayStocks(water, milk, beans, cups, money *int) {
	fmt.Printf(`The coffee machine has:
%d ml of water
%d ml of milk
%d g of coffee beans
%d disposable cups
$%d of money

`, *water, *milk, *beans, *cups, *money)
}

func displayMenu() string {
	var action string
	fmt.Println("Write action (buy, fill, take, remaining, exit):")
	fmt.Scan(&action)
	return action
}

func startCoffeeMaking(water, milk, beans, cups, money *int) {
	for {
		action := displayMenu()
		switch action {
		case "buy":
			buy(&water, &milk, &beans, &cups, &money)
		case "fill":
			fill(&water, &milk, &beans, &cups)
		case "take":
			take(&money)
		case "exit":
			return
		case "remaining":
			displayStocks(water, milk, beans, cups, money)
		}
	}
}

func main() {
	water, milk, beans, cups, money := 400, 540, 120, 9, 550
	startCoffeeMaking(&water, &milk, &beans, &cups, &money)
}
