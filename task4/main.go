package main

import (
	"fmt"
	"log"
)

func main() {
	// Feel free to use the main function for testing your functions
	c := Price.String(2595)
	fmt.Println(c)

	RegisterItem(Prices, "meat", 400)
	RegisterItem(Prices, "fish", 500)
	RegisterItem(Prices, "milk", 300)

	keys := []string{"milk", "bread"}
	var totalPrice Price = 0

	for _, i := range keys {
		totalPrice += Prices[i]
	}
	k := &Cart{Items: keys, TotalPrice: totalPrice}
	fmt.Println(k.hasMilk())
	fmt.Println(k.HasItem("meat"))
	k.AddItem("meat")
	k.AddItem("peanut butter")
	k.Checkout()
}

// Price is the cost of something in US cents.
type Price int64

// String is the string representation of a Price
// These should be represented in US Dollars
// Example: 2595 cents => $25.95
func (p Price) String() string {
	// TODO
	k := fmt.Sprintf("%.2f", float64(p)/100)
	return " $" + k
}

// Prices is a map from an item to its price.
var Prices = map[string]Price{
	"eggs":          219,
	"bread":         199,
	"milk":          295,
	"peanut butter": 445,
	"chocolate":     150,
}

// RegisterItem adds the new item in the prices map.
// If the item is already in the prices map, a warning should be displayed to the user,
// but the value should be overwritten.
// Bonus (1pt) - Use the "log" package to print the error to the user
func RegisterItem(prices map[string]Price, item string, price Price) {
	for z := range prices {
		if z == item {
			prices[item] = price
			log.Printf("We already have %v", item)
		}
	}
	prices[item] = price
	// TODO
}

// Cart is a struct representing a shopping cart of items.
type Cart struct {
	Items      []string
	TotalPrice Price
}

// hasMilk returns whether the shopping cart has "milk".
func (c *Cart) hasMilk() bool {
	// TODO
	for _, item := range c.Items {
		if item == "milk" {
			return true
		}
	}
	return false
}

// HasItem returns whether the shopping cart has the provided item name.
func (c *Cart) HasItem(item string) bool {
	for _, i := range c.Items {
		if i == item {
			return true
		}
	}
	// TODO
	return false
}

// AddItem adds the provided item to the cart and update the cart balance.
// If item is not found in the prices map, then do not add it and print an error.
// Bonus (1pt) - Use the "log" package to print the error to the user
func (c *Cart) AddItem(item string) {
	true_or_false := false
	for i := range Prices {
		if i == item {
			true_or_false = true
		}
	}
	if !true_or_false {
		error_message := "You made a misstake with item"
		log.Println(error_message)
		return
	}
	c.Items = append(c.Items, item)
	c.TotalPrice += Prices[item]
	fmt.Println(c.Items)
	fmt.Println(c.TotalPrice)
	// TODO
}

// Checkout displays the final cart balance and clears the cart completely.
func (c *Cart) Checkout() {
	items := ""
	for _, i := range c.Items {
		items += i + ", "
	}
	fmt.Printf("You have bought: %sand you should pay%v", items, c.TotalPrice)
	// TODO
}
