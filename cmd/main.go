package main

import (
	"facade/pkg"
	"fmt"
)

var (
	bank = &pkg.Bank{
		Name:  "My Bank",
		Cards: []pkg.Card{},
	}
	card1 = &pkg.Card{
		Name:    "CRD-1",
		Balance: 100,
		Bank:    bank,
	}
	card2 = &pkg.Card{
		Name:    "CRD-2",
		Balance: 200,
		Bank:    bank,
	}
	user1 = &pkg.User{
		Name: "John",
		Card: card1,
	}
	user2 = &pkg.User{
		Name: "Jane",
		Card: card2,
	}
	product = &pkg.Product{
		Name:  "Laptop",
		Price: 150,
	}
	shop = &pkg.Shop{
		Name:     "My Shop",
		Products: []pkg.Product{*product},
	}
)

func main() {
	fmt.Println("Welcome to the shop!")
	fmt.Println("Add cards")
	bank.Cards = append(bank.Cards, *card1, *card2)
	fmt.Printf("[%s]\n", product.Name)
	err := shop.Buy(product, user1)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = shop.Buy(product, user2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Thank you for shopping!")
}
